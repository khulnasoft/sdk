// Copyright 2016-2020, KhulnaSoft, Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auto

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/blang/semver"
	"github.com/khulnasoft/sdk"
	"github.com/khulnasoft/sdk/go/common/env"
	"github.com/khulnasoft/sdk/go/common/slice"
	"github.com/khulnasoft/sdk/go/common/util/httputil"
)

const unknownErrorCode = -2

// KhulnasoftCommand manages the Khulnasoft CLI and runs operations.
type KhulnasoftCommand interface {
	Run(ctx context.Context,
		workdir string,
		stdin io.Reader,
		additionalOutput []io.Writer,
		additionalErrorOutput []io.Writer,
		additionalEnv []string,
		args ...string) (string, string, int, error)
	Version() semver.Version
}

type KhulnasoftCommandOptions struct {
	// Version is the version to install or validate.
	Version semver.Version
	// Root sets the directory where the CLI should be installed to, or from
	// where the CLI should be retrieved.
	Root string
	// SkipVersionCheck is used to disable the validation of the found Khulnasoft
	// binary.
	SkipVersionCheck bool
}

// withDefaults returns a new copy of the options with default values set.
// Version defaults to the CLI version matching the current SDK release. Root
// defaults to $HOME/.khulnasoft/versions/$VERSION.
func (opts *KhulnasoftCommandOptions) withDefaults() (*KhulnasoftCommandOptions, error) {
	newOpts := &KhulnasoftCommandOptions{
		Version:          opts.Version,
		Root:             opts.Root,
		SkipVersionCheck: opts.SkipVersionCheck,
	}
	if newOpts.Version.EQ(semver.Version{}) {
		newOpts.Version = sdk.Version
	}

	if newOpts.Root == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		newOpts.Root = filepath.Join(home, ".khulnasoft", "versions", newOpts.Version.String())
	}

	return newOpts, nil
}

type khulnasoftCommand struct {
	version semver.Version
	command string
}

// NewKhulnasoftCommand creates a Khulnasoft instance that uses the installation in
// `opts.Root`.  Defaults to using the Khulnasoft binary found in $PATH if no
// installation root is specified.  If `opts.Version` is specified, it
// validates that the CLI is compatible with the requested version and throws
// an error if not. This validation can be skipped by setting the environment
// variable `KHULNASOFT_AUTOMATION_API_SKIP_VERSION_CHECK` or setting
// `opts.SkipVersionCheck` to `true`. Note that the environment variable
// always takes precedence. If it is set it is not possible to re-enable the
// validation with `opts.SkipVersionCheck`.
func NewKhulnasoftCommand(opts *KhulnasoftCommandOptions) (KhulnasoftCommand, error) {
	if opts == nil {
		opts = &KhulnasoftCommandOptions{}
	}
	command := "khulnasoft"
	if opts.Root != "" {
		command = filepath.Join(opts.Root, "bin", "khulnasoft")
		if runtime.GOOS == "windows" {
			command += ".exe"
		}
	}

	cmd := exec.Command(command, "version")
	out, err := cmd.Output()
	if err != nil {
		return khulnasoftCommand{}, fmt.Errorf("failed to run `khulnasoft version`: %w", err)
	}
	currentVersion := strings.TrimSpace(string(out))
	minimum := minimumVersion
	if opts.Version.GT(minimum) {
		minimum = opts.Version
	}
	skipVersionCheck := opts.SkipVersionCheck || env.SkipVersionCheck.Value()
	version, err := parseAndValidateKhulnasoftVersion(minimum, currentVersion, skipVersionCheck)
	if err != nil {
		return khulnasoftCommand{}, err
	}

	return khulnasoftCommand{
		version: version,
		command: command,
	}, nil
}

// InstallKhulnasoftCommand downloads and installs the Khulnasoft CLI. By default the
// CLI version matching the current SDK release is installed in
// $HOME/.khulnasoft/versions/$VERSION. Set `opts.Root` to specify a different
// directory, and `opts.Version` to install a custom version.
func InstallKhulnasoftCommand(ctx context.Context, opts *KhulnasoftCommandOptions) (
	KhulnasoftCommand,
	error,
) {
	if opts == nil {
		opts = &KhulnasoftCommandOptions{}
	}
	opts, err := opts.withDefaults()
	if err != nil {
		return khulnasoftCommand{}, err
	}
	khulnasoft, err := NewKhulnasoftCommand(opts)
	if err == nil {
		// Found an installation and it satisfies the version requirement
		return khulnasoft, nil
	}
	if runtime.GOOS == "windows" {
		if err := installWindows(ctx, opts.Version, opts.Root); err != nil {
			return khulnasoftCommand{}, err
		}
	} else {
		if err := installPosix(ctx, opts.Version, opts.Root); err != nil {
			return khulnasoftCommand{}, err
		}
	}
	return NewKhulnasoftCommand(opts)
}

func downloadToTmpFile(ctx context.Context, url, filePattern string) (_ string, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}
	resp, err := httputil.DoWithRetry(req, http.DefaultClient)
	if err != nil {
		return "", fmt.Errorf("failed to download %s: %w", url, err)
	}
	defer resp.Body.Close()
	scriptData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response from %s: %w", url, err)
	}
	tmp, err := os.CreateTemp("", filePattern)
	if err != nil {
		return "", err
	}
	scriptPath := tmp.Name()
	defer func() {
		if err != nil {
			os.Remove(scriptPath)
		}
	}()
	// The permissions here are ignored because the tmp file already exists.
	// We need to explicitly call chmod below to set the desired permissions.
	err = os.WriteFile(scriptPath, scriptData, 0o600)
	if err != nil {
		return "", err
	}
	if err = tmp.Close(); err != nil {
		return "", err
	}
	err = os.Chmod(scriptPath, 0o700)
	if err != nil {
		return "", err
	}
	return scriptPath, nil
}

func installWindows(ctx context.Context, version semver.Version, root string) error {
	scriptPath, err := downloadToTmpFile(ctx, "https://get.khulnasoft.com/install.ps1", "install-*.ps1")
	if err != nil {
		return err
	}
	defer os.Remove(scriptPath)
	command := "powershell.exe"
	if os.Getenv("SystemRoot") != "" {
		command = filepath.Join(os.Getenv("SystemRoot"), "System32", "WindowsPowerShell", "v1.0", "powershell.exe")
	}
	args := []string{
		"-NoProfile",
		"-InputFormat",
		"None",
		"-ExecutionPolicy",
		"Bypass",
		"-File",
		scriptPath,
		"-NoEditPath",
		"-InstallRoot",
		root,
		"-Version",
		version.String(),
	}
	cmd := exec.CommandContext(ctx, command, args...)
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("installation failed with %w: %s", err, out)
	}
	return nil
}

func installPosix(ctx context.Context, version semver.Version, root string) error {
	scriptPath, err := downloadToTmpFile(ctx, "https://get.khulnasoft.com/install.sh", "install-*.sh")
	if err != nil {
		return err
	}
	defer os.Remove(scriptPath)
	args := []string{
		"--no-edit-path",
		"--install-root", root,
		"--version", version.String(),
	}
	cmd := exec.CommandContext(ctx, scriptPath, args...)
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("installation failed with %w: %s", err, out)
	}
	return nil
}

// Run executes a Khulnasoft CLI command.
func (p khulnasoftCommand) Run(ctx context.Context,
	workdir string,
	stdin io.Reader,
	additionalOutput []io.Writer,
	additionalErrorOutput []io.Writer,
	additionalEnv []string,
	args ...string,
) (string, string, int, error) {
	// all commands should be run in non-interactive mode.
	// this causes commands to fail rather than prompting for input (and thus hanging indefinitely)
	args = withNonInteractiveArg(args)
	cmd := exec.CommandContext(ctx, p.command, args...) //nolint:gosec
	cmd.Dir = workdir
	env := append(os.Environ(), additionalEnv...)
	if filepath.IsAbs(p.command) {
		khulnasoftBin := filepath.Dir(p.command)
		env = fixupPath(env, khulnasoftBin)
	}
	cmd.Env = env

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	additionalOutput = append(additionalOutput, &stdout)
	additionalErrorOutput = append(additionalErrorOutput, &stderr)
	cmd.Stdout = io.MultiWriter(additionalOutput...)
	cmd.Stderr = io.MultiWriter(additionalErrorOutput...)
	cmd.Stdin = stdin

	code := unknownErrorCode
	err := cmd.Run()
	if exitError, ok := err.(*exec.ExitError); ok {
		code = exitError.ExitCode()
	} else if err == nil {
		// If there was no error then the exit code was 0
		code = 0
	}
	return stdout.String(), stderr.String(), code, err
}

// The version of the current Khulnasoft CLI installation.
func (p khulnasoftCommand) Version() semver.Version {
	return p.version
}

func withNonInteractiveArg(args []string) []string {
	out := slice.Prealloc[string](len(args))
	seen := false
	for _, a := range args {
		out = append(out, a)
		if a == "--non-interactive" {
			seen = true
		}
	}
	if !seen {
		out = append(out, "--non-interactive")
	}
	return out
}

// fixupPath returns a new copy of `env` where the `PATH` entry contains
// khulnasoftBin as the first path.
func fixupPath(env []string, khulnasoftBin string) []string {
	newEnv := make([]string, len(env))
	copy(newEnv, env)
	pathIndex := -1
	for i, e := range env {
		// Case-insensitive compare, as Windows will normally be "Path", not "PATH".
		if len(e) >= 5 && strings.EqualFold(e[0:5], "PATH=") {
			pathIndex = i
			break
		}
	}
	if pathIndex >= 0 {
		pathEntry := khulnasoftBin
		oldPath := env[pathIndex][5:]
		if oldPath != "" {
			pathEntry = khulnasoftBin + string(os.PathListSeparator) + oldPath
		}
		newEnv[pathIndex] = "PATH=" + pathEntry
	} else {
		newEnv = append(newEnv, "PATH="+khulnasoftBin)
	}
	return newEnv
}

//nolint:lll
func parseAndValidateKhulnasoftVersion(minVersion semver.Version, currentVersion string, optOut bool) (semver.Version, error) {
	version, err := semver.ParseTolerant(currentVersion)
	if err != nil && !optOut {
		return semver.Version{}, fmt.Errorf("Unable to parse Khulnasoft CLI version (skip with %s=true): %w", env.SkipVersionCheck.Var().Name(), err)
	}
	if optOut {
		return version, nil
	}
	if minVersion.Major < version.Major {
		return semver.Version{}, fmt.Errorf("Major version mismatch. You are using Khulnasoft CLI version %s with Automation SDK v%v. Please update the SDK.", currentVersion, minVersion.Major) //nolint
	}
	if minVersion.GT(version) {
		return semver.Version{}, fmt.Errorf("Minimum version requirement failed. The minimum CLI version requirement is %s, your current CLI version is %s. Please update the Khulnasoft CLI.", minimumVersion, currentVersion) //nolint
	}
	return version, nil
}
