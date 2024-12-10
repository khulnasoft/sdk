// Copyright 2016-2023, KhulnaSoft, Ltd.
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

// khulnasoft-language-python serves as the "language host" for Khulnasoft programs written in Python.  It is ultimately
// responsible for spawning the language runtime that executes the program.
//
// The program being executed is executed by a shim script called `khulnasoft-language-python-exec`. This script is
// written in the hosted language (in this case, Python) and is responsible for initiating RPC links to the resource
// monitor and engine.
//
// It's therefore the responsibility of this program to implement the LanguageHostServer endpoint by spawning
// instances of `khulnasoft-language-python-exec` and forwarding the RPC request arguments to the command-line.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unicode"

	"github.com/blang/semver"
	"github.com/khulnasoft/sdk/go/common/resource/plugin"
	"github.com/khulnasoft/sdk/go/common/slice"
	"github.com/khulnasoft/sdk/go/common/util/cmdutil"
	"github.com/khulnasoft/sdk/go/common/util/contract"
	"github.com/khulnasoft/sdk/go/common/util/fsutil"
	"github.com/khulnasoft/sdk/go/common/util/logging"
	"github.com/khulnasoft/sdk/go/common/util/rpcutil"
	"github.com/khulnasoft/sdk/go/common/version"
	"github.com/khulnasoft/sdk/go/common/workspace"
	"github.com/khulnasoft/sdk/go/khulnasoft-internal/netutil"
	khulnasoftrpc "github.com/khulnasoft/sdk/proto/go"
	"github.com/khulnasoft/sdk/python/toolchain"
	"github.com/nxadm/tail"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"

	hclsyntax "github.com/khulnasoft/khulnasoft/pkg/v3/codegen/hcl2/syntax"
	"github.com/khulnasoft/khulnasoft/pkg/v3/codegen/pcl"
	codegen "github.com/khulnasoft/khulnasoft/pkg/v3/codegen/python"
	"github.com/khulnasoft/khulnasoft/pkg/v3/codegen/schema"
)

const (
	// By convention, the executor is the name of the current program (khulnasoft-language-python) plus this suffix.
	pythonDefaultExec = "khulnasoft-language-python-exec" // the exec shim for Khulnasoft to run Python programs.

	// The runtime expects the config object to be saved to this environment variable.
	khulnasoftConfigVar = "KHULNASOFT_CONFIG"

	// The runtime expects the array of secret config keys to be saved to this environment variable.
	//nolint:gosec
	khulnasoftConfigSecretKeysVar = "KHULNASOFT_CONFIG_SECRET_KEYS"

	// A exit-code we recognize when the python process exits.  If we see this error, there's no
	// need for us to print any additional error messages since the user already got a a good
	// one they can handle.
	pythonProcessExitedAfterShowingUserActionableMessage = 32

	// The preferred debug port.  Chosen arbitrarily.
	preferredDebugPort = 58791
)

var (
	// The minimum python version that Khulnasoft supports
	minimumSupportedPythonVersion = semver.MustParse("3.7.0")
	// Any version less then `eolPythonVersion` is EOL.
	eolPythonVersion = semver.MustParse("3.7.0")
	// An url to the issue discussing EOL.
	eolPythonVersionIssue = "https://github.com/khulnasoft/khulnasoft/issues/8131"
)

// Launches the language host RPC endpoint, which in turn fires up an RPC server implementing the
// LanguageRuntimeServer RPC endpoint.
func main() {
	var tracing string
	flag.StringVar(&tracing, "tracing", "", "Emit tracing to a Zipkin-compatible tracing endpoint")
	flag.String("virtualenv", "", "[obsolete] Virtual environment path to use")
	flag.String("root", "", "[obsolete] Project root path to use")
	flag.String("typechecker", "", "[obsolete] Use a typechecker to type check")
	flag.String("toolchain", "pip", "[obsolete] Select the package manager to use for dependency management.")

	// You can use the below flag to request that the language host load a specific executor instead of probing the
	// PATH.  This can be used during testing to override the default location.
	var givenExecutor string
	flag.StringVar(&givenExecutor, "use-executor", "",
		"Use the given program as the executor instead of looking for one on PATH")

	flag.Parse()
	args := flag.Args()
	logging.InitLogging(false, 0, false)
	cmdutil.InitTracing("khulnasoft-language-python", "khulnasoft-language-python", tracing)

	var pythonExec string
	if givenExecutor == "" {
		// By default, the -exec script is installed next to the language host.
		thisPath, err := os.Executable()
		if err != nil {
			err = fmt.Errorf("could not determine current executable: %w", err)
			cmdutil.Exit(err)
		}

		pathExec := filepath.Join(filepath.Dir(thisPath), pythonDefaultExec)
		if _, err = os.Stat(pathExec); os.IsNotExist(err) {
			err = fmt.Errorf("missing executor %s: %w", pathExec, err)
			cmdutil.Exit(err)
		}

		logging.V(3).Infof("language host identified executor from path: `%s`", pathExec)
		pythonExec = pathExec
	} else {
		logging.V(3).Infof("language host asked to use specific executor: `%s`", givenExecutor)
		pythonExec = givenExecutor
	}

	// Optionally pluck out the engine so we can do logging, etc.
	var engineAddress string
	if len(args) > 0 {
		engineAddress = args[0]
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// map the context Done channel to the rpcutil boolean cancel channel
	cancelChannel := make(chan bool)
	go func() {
		<-ctx.Done()
		cancel() // deregister signal handler
		close(cancelChannel)
	}()
	err := rpcutil.Healthcheck(ctx, engineAddress, 5*time.Minute, cancel)
	if err != nil {
		cmdutil.Exit(fmt.Errorf("could not start health check host RPC server: %w", err))
	}

	// Fire up a gRPC server, letting the kernel choose a free port.
	handle, err := rpcutil.ServeWithOptions(rpcutil.ServeOptions{
		Cancel: cancelChannel,
		Init: func(srv *grpc.Server) error {
			host := newLanguageHost(pythonExec, engineAddress, tracing, "")
			khulnasoftrpc.RegisterLanguageRuntimeServer(srv, host)
			return nil
		},
		Options: rpcutil.OpenTracingServerInterceptorOptions(nil),
	})
	if err != nil {
		cmdutil.Exit(fmt.Errorf("could not start language host RPC server: %w", err))
	}

	// Otherwise, print out the port so that the spawner knows how to reach us.
	fmt.Printf("%d\n", handle.Port)

	// And finally wait for the server to stop serving.
	if err := <-handle.Done; err != nil {
		cmdutil.Exit(fmt.Errorf("language host RPC stopped serving: %w", err))
	}
}

// pythonLanguageHost implements the LanguageRuntimeServer interface
// for use as an API endpoint.
type pythonLanguageHost struct {
	khulnasoftrpc.UnsafeLanguageRuntimeServer

	exec          string
	engineAddress string
	tracing       string

	// This is used by conformance testing to set the typechecker to use in ProgramGen.
	typechecker string
}

func parseOptions(root string, programDir string, options map[string]interface{}) (toolchain.PythonOptions, error) {
	pythonOptions := toolchain.PythonOptions{
		Root:       root,
		ProgramDir: programDir,
	}

	if virtualenv, ok := options["virtualenv"]; ok {
		if virtualenv, ok := virtualenv.(string); ok {
			pythonOptions.Virtualenv = virtualenv
		} else {
			return pythonOptions, errors.New("virtualenv option must be a string")
		}
	}

	if typechecker, ok := options["typechecker"]; ok {
		if typechecker, ok := typechecker.(string); ok {
			switch typechecker {
			case "mypy":
				pythonOptions.Typechecker = toolchain.TypeCheckerMypy
			case "pyright":
				pythonOptions.Typechecker = toolchain.TypeCheckerPyright
			default:
				return pythonOptions, fmt.Errorf("unsupported typechecker option: %s", typechecker)
			}
		} else {
			return pythonOptions, errors.New("typechecker option must be a string")
		}
	}

	if tc, ok := options["toolchain"]; ok {
		if tc, ok := tc.(string); ok {
			switch tc {
			case "pip":
				pythonOptions.Toolchain = toolchain.Pip
			case "poetry":
				pythonOptions.Toolchain = toolchain.Poetry
			case "uv":
				pythonOptions.Toolchain = toolchain.Uv
			default:
				return pythonOptions, fmt.Errorf("unsupported toolchain option: %s", tc)
			}
		} else {
			return pythonOptions, errors.New("toolchain option must be a string")
		}
	}

	return pythonOptions, nil
}

func newLanguageHost(exec, engineAddress, tracing, typechecker string,
) khulnasoftrpc.LanguageRuntimeServer {
	return &pythonLanguageHost{
		exec:          exec,
		engineAddress: engineAddress,
		tracing:       tracing,
		typechecker:   typechecker,
	}
}

func (host *pythonLanguageHost) connectToEngine() (khulnasoftrpc.EngineClient, io.Closer, error) {
	conn, err := grpc.NewClient(
		host.engineAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		rpcutil.GrpcChannelOptions(),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("language host could not make connection to engine: %w", err)
	}

	engineClient := khulnasoftrpc.NewEngineClient(conn)
	return engineClient, conn, nil
}

func (host *pythonLanguageHost) GetRequiredPackages(ctx context.Context,
	req *khulnasoftrpc.GetRequiredPackagesRequest,
) (*khulnasoftrpc.GetRequiredPackagesResponse, error) {
	opts, err := parseOptions(req.Info.RootDirectory, req.Info.ProgramDirectory, req.Info.Options.AsMap())
	if err != nil {
		return nil, err
	}

	tc, err := toolchain.ResolveToolchain(opts)
	if err != nil {
		return nil, err
	}

	stdout, stderr, err := host.createEngineWriters(ctx)
	if err != nil {
		return nil, err
	}
	if err := tc.EnsureVenv(ctx, req.Info.ProgramDirectory, false, /*useLanguageVersionTools */
		true /* showOutput */, stdout, stderr); err != nil {
		return nil, err
	}

	validateVersion(ctx, opts)

	// Now, determine which Khulnasoft packages are installed.
	khulnasoftPackages, err := determineKhulnasoftPackages(ctx, opts)
	if err != nil {
		return nil, err
	}

	packages := []*khulnasoftrpc.PackageDependency{}
	for _, pkg := range khulnasoftPackages {
		pkg, err := determinePackageDependency(pkg)
		if err != nil {
			return nil, err
		}

		if pkg != nil {
			packages = append(packages, pkg)
		}
	}

	return &khulnasoftrpc.GetRequiredPackagesResponse{Packages: packages}, nil
}

// GetRequiredPlugins computes the complete set of anticipated plugins required by a program.
func (host *pythonLanguageHost) GetRequiredPlugins(ctx context.Context,
	req *khulnasoftrpc.GetRequiredPluginsRequest,
) (*khulnasoftrpc.GetRequiredPluginsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequiredPlugins not implemented")
}

func (host *pythonLanguageHost) Pack(ctx context.Context, req *khulnasoftrpc.PackRequest) (*khulnasoftrpc.PackResponse, error) {
	tc, err := toolchain.ResolveToolchain(toolchain.PythonOptions{
		Toolchain: toolchain.Pip,
	})
	if err != nil {
		return nil, err
	}
	// ensure build is up-to-date
	buildUpgradeCmd, err := tc.ModuleCommand(ctx, "pip", "install", "--upgrade", "build")
	if err != nil {
		return nil, err
	}
	buildUpgradeCmd.Stdout = os.Stdout
	buildUpgradeCmd.Stderr = os.Stderr
	err = buildUpgradeCmd.Run()
	if err != nil {
		return nil, fmt.Errorf("install build tools: %w", err)
	}

	tmp, err := os.MkdirTemp("", "khulnasoft-python-pack")
	if err != nil {
		return nil, fmt.Errorf("create temporary directory: %w", err)
	}

	buildCmd, err := tc.ModuleCommand(ctx, "build", "--outdir", tmp)
	if err != nil {
		return nil, err
	}
	buildCmd.Dir = req.PackageDirectory

	var stdout, stderr bytes.Buffer
	buildCmd.Stdout = &stdout
	buildCmd.Stderr = &stderr

	err = buildCmd.Run()
	logging.V(5).Infof("Pack stdout: %s", stdout.String())
	logging.V(5).Infof("Pack stderr: %s", stderr.String())
	if err != nil {
		return nil, fmt.Errorf("run python build: %w\n%s\n%s", err, stdout.String(), stderr.String())
	}

	// prefer .whl but return .tar.gz if no .whl is found
	files, err := os.ReadDir(tmp)
	if err != nil {
		return nil, fmt.Errorf("read temporary directory: %w", err)
	}

	var found string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".whl") {
			found = file.Name()
			break
		}
		if strings.HasSuffix(file.Name(), ".tar.gz") {
			found = file.Name()
		}
	}

	// Copy the found file to the destination directory
	if found == "" {
		return nil, fmt.Errorf("no .whl or .tar.gz file found\n%s", stderr.String())
	}

	src := filepath.Join(tmp, found)
	dst := filepath.Join(req.DestinationDirectory, found)
	err = fsutil.CopyFile(dst, src, nil)
	if err != nil {
		return nil, fmt.Errorf("copy file: %w", err)
	}

	return &khulnasoftrpc.PackResponse{
		ArtifactPath: dst,
	}, nil
}

// createEngineWriters creates a pair of writers that can be used to log messages to the engine.
func (host *pythonLanguageHost) createEngineWriters(ctx context.Context) (io.Writer, io.Writer, error) {
	// Make a connection to the real engine that we will log messages to.
	conn, err := grpc.NewClient(
		host.engineAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		rpcutil.GrpcChannelOptions(),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("language host could not make connection to engine: %w", err)
	}

	// Make a client around that connection.
	engineClient := khulnasoftrpc.NewEngineClient(conn)

	// Create writers that log the output of the install operation as ephemeral messages.
	streamID := rand.Int31() //nolint:gosec

	infoWriter := &logWriter{
		ctx:          ctx,
		engineClient: engineClient,
		streamID:     streamID,
		severity:     khulnasoftrpc.LogSeverity_INFO,
	}

	errorWriter := &logWriter{
		ctx:          ctx,
		engineClient: engineClient,
		streamID:     streamID,
		severity:     khulnasoftrpc.LogSeverity_ERROR,
	}

	return infoWriter, errorWriter, nil
}

type logWriter struct {
	ctx          context.Context
	engineClient khulnasoftrpc.EngineClient
	streamID     int32
	severity     khulnasoftrpc.LogSeverity
}

func (w *logWriter) Write(p []byte) (n int, err error) {
	val := string(p)
	if _, err := w.engineClient.Log(w.ctx, &khulnasoftrpc.LogRequest{
		Message:   strings.ToValidUTF8(val, "�"),
		Urn:       "",
		Ephemeral: true,
		StreamId:  w.streamID,
		Severity:  w.severity,
	}); err != nil {
		return 0, err
	}
	return len(val), nil
}

// These packages are known not to have any plugins.
// TODO[khulnasoft/khulnasoft#5863]: Remove this once the `khulnasoft-policy` package includes a `khulnasoft-plugin.json`
// file that indicates the package does not have an associated plugin, and enough time has passed.
var packagesWithoutPlugins = map[string]struct{}{
	// We include both the hyphen and underscore variants of the package name
	// to account for the fact that later versions of the package will come
	// back from `python -m pip list` as the underscore variant due to a
	// behavior change in setuptools where it keeps underscores rather than
	// replacing them with hyphens.
	"khulnasoft-policy": {},
	"khulnasoft_policy": {},
}

// Returns if pkg is a khulnasoft package.
//
// We check:
// 1. If there is a khulnasoft-plugin.json file.
// 2. If the first segment is "khulnasoft". This implies a first party package.
func isKhulnasoftPackage(pkg toolchain.PythonPackage) bool {
	plugin, err := readKhulnasoftPluginJSON(pkg)
	if err == nil && plugin != nil {
		return true
	}

	return strings.HasPrefix(pkg.Name, "khulnasoft_") || strings.HasPrefix(pkg.Name, "khulnasoft-")
}

func readKhulnasoftPluginJSON(pkg toolchain.PythonPackage) (*plugin.KhulnasoftPluginJSON, error) {
	// The name of the module inside the package can be different from the package name.
	// However, our convention is to always use the same name, e.g. a package name of
	// "khulnasoft-aws" will have a module named "khulnasoft_aws", so we can determine the module
	// by replacing hyphens with underscores.
	packageModuleName := strings.ReplaceAll(pkg.Name, "-", "_")
	khulnasoftPluginFilePath := filepath.Join(pkg.Location, packageModuleName, "khulnasoft-plugin.json")
	logging.V(5).Infof("readKhulnasoftPluginJSON: khulnasoft-plugin.json file path: %s", khulnasoftPluginFilePath)

	plugin, err := plugin.LoadKhulnasoftPluginJSON(khulnasoftPluginFilePath)
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return plugin, nil
}

func determineKhulnasoftPackages(ctx context.Context, options toolchain.PythonOptions) ([]toolchain.PythonPackage, error) {
	logging.V(5).Infof("GetRequiredPlugins: Determining khulnasoft packages")

	tc, err := toolchain.ResolveToolchain(options)
	if err != nil {
		return nil, err
	}
	packages, err := tc.ListPackages(ctx, true /* transitive */)
	if err != nil {
		return nil, err
	}

	// Only return Khulnasoft packages.
	khulnasoftPackages := slice.Prealloc[toolchain.PythonPackage](len(packages))
	for _, pkg := range packages {
		if !isKhulnasoftPackage(pkg) {
			continue
		}

		// Skip packages that are known not have an associated plugin.
		if _, ok := packagesWithoutPlugins[pkg.Name]; ok {
			continue
		}

		khulnasoftPackages = append(khulnasoftPackages, pkg)
	}

	logging.V(5).Infof("GetRequiredPlugins: Khulnasoft packages: %#v", khulnasoftPackages)

	return khulnasoftPackages, nil
}

// determinePackageDependency attempts to determine a khulnasoft package associated with a python package. It
// checks to see if the package contains a khulnasoft-plugin.json file and uses the information in that file to
// determine the plugin. If `resource` in khulnasoft-plugin.json is set to false, nil is returned. If the name or
// version aren't specified in the file, these values are derived from the package name and version. If the
// plugin version cannot be determined from the package version, nil is returned.
func determinePackageDependency(pkg toolchain.PythonPackage) (*khulnasoftrpc.PackageDependency, error) {
	var name, version, server string
	var parameterization *khulnasoftrpc.PackageParameterization
	plugin, err := readKhulnasoftPluginJSON(pkg)
	if plugin != nil && err == nil {
		// If `resource` is set to false, the Khulnasoft package has indicated that there is no associated plugin.
		// Ignore it.
		if !plugin.Resource {
			logging.V(5).Infof("GetRequiredPlugins: Ignoring package %s with resource set to false", pkg.Name)
			return nil, nil
		}

		if plugin.Parameterization != nil {
			parameterization = &khulnasoftrpc.PackageParameterization{
				Name:    plugin.Parameterization.Name,
				Version: plugin.Parameterization.Version,
				Value:   plugin.Parameterization.Value,
			}
		}

		name, version, server = plugin.Name, plugin.Version, plugin.Server
	} else if err != nil {
		logging.V(5).Infof("GetRequiredPlugins: err: %v", err)
		return nil, err
	}

	if name == "" {
		name = strings.TrimPrefix(pkg.Name, "khulnasoft-")
	}

	if version == "" {
		// The packageVersion may include additional pre-release tags (e.g. "2.14.0a1605583329" for an alpha
		// release, "2.14.0b1605583329" for a beta release, "2.14.0rc1605583329" for an rc release, etc.).
		// Unfortunately, this is not enough information to determine the plugin version. A package version of
		// "3.31.0a1605189729" will have an associated plugin with a version of "3.31.0-alpha.1605189729+42435656".
		// The "+42435656" suffix cannot be determined so the plugin version cannot be determined. In such cases,
		// log the issue and skip the package.
		version, err = determinePluginVersion(pkg.Version)
		if err != nil {
			logging.V(5).Infof(
				"GetRequiredPlugins: Could not determine plugin version for package %s with version %s: %s",
				pkg.Name, pkg.Version, err)
			return nil, nil
		}
	}
	if !strings.HasPrefix(version, "v") {
		// Add "v" prefix if not already present.
		version = "v" + version
	}

	result := &khulnasoftrpc.PackageDependency{
		Name:             name,
		Version:          version,
		Kind:             "resource",
		Server:           server,
		Parameterization: parameterization,
	}

	logging.V(5).Infof("GetRequiredPlugins: Determining plugin dependency: %#v", result)
	return result, nil
}

// determinePluginVersion attempts to convert a PEP440 package version into a plugin version.
//
// Supported versions:
//
//	PEP440 defines a version as `[N!]N(.N)*[{a|b|rc}N][.postN][.devN]`, but
//	determinePluginVersion only supports a subset of that. Translations are provided for
//	`N(.N)*[{a|b|rc}N][.postN][.devN]`.
//
// Translations:
//
//	We ensure that there are at least 3 version segments. Missing segments are `0`
//	padded.
//	Example: 1.0 => 1.0.0
//
//	We translate a,b,rc to alpha,beta,rc respectively with a hyphen separator.
//	Example: 1.2.3a4 => 1.2.3-alpha.4, 1.2.3rc4 => 1.2.3-rc.4
//
//	We translate `.post` and `.dev` by replacing the `.` with a `+`. If both `.post`
//	and `.dev` are present, only one separator is used.
//	Example: 1.2.3.post4 => 1.2.3+post4, 1.2.3.post4.dev5 => 1.2.3+post4dev5
//
// Reference on PEP440: https://www.python.org/dev/peps/pep-0440/
func determinePluginVersion(packageVersion string) (string, error) {
	if len(packageVersion) == 0 {
		return "", errors.New("cannot parse empty string")
	}
	// Verify ASCII
	for i := 0; i < len(packageVersion); i++ {
		c := packageVersion[i]
		if c > unicode.MaxASCII {
			return "", fmt.Errorf("byte %d is not ascii", i)
		}
	}

	parseNumber := func(s string) (string, string) {
		i := 0
		for _, c := range s {
			if c > '9' || c < '0' {
				break
			}
			i++
		}
		return s[:i], s[i:]
	}

	// Explicitly err on epochs
	if num, maybeEpoch := parseNumber(packageVersion); num != "" && strings.HasPrefix(maybeEpoch, "!") {
		return "", errors.New("epochs are not supported")
	}

	segments := []string{}
	var num string
	rest := packageVersion
	foundDot := false
	for {
		if num, rest = parseNumber(rest); num != "" {
			foundDot = false
			segments = append(segments, num)
			if strings.HasPrefix(rest, ".") {
				rest = rest[1:]
				foundDot = true
			} else {
				break
			}
		} else {
			break
		}
	}
	if foundDot {
		rest = "." + rest
	}

	for len(segments) < 3 {
		segments = append(segments, "0")
	}

	if rest == "" {
		r := strings.Join(segments, ".")
		return r, nil
	}

	var preRelease string

	switch {
	case rest[0] == 'a':
		preRelease, rest = parseNumber(rest[1:])
		preRelease = "-alpha." + preRelease
	case rest[0] == 'b':
		preRelease, rest = parseNumber(rest[1:])
		preRelease = "-beta." + preRelease
	case strings.HasPrefix(rest, "rc"):
		preRelease, rest = parseNumber(rest[2:])
		preRelease = "-rc." + preRelease
	}

	var postRelease string
	if strings.HasPrefix(rest, ".post") {
		postRelease, rest = parseNumber(rest[5:])
		postRelease = "+post" + postRelease
	}

	var developmentRelease string
	if strings.HasPrefix(rest, ".dev") {
		developmentRelease, rest = parseNumber(rest[4:])
		join := ""
		if postRelease == "" {
			join = "+"
		}
		developmentRelease = join + "dev" + developmentRelease
	}

	if rest != "" {
		return "", fmt.Errorf("'%s' still unparsed", rest)
	}

	result := strings.Join(segments, ".") + preRelease + postRelease + developmentRelease

	return result, nil
}

// debugCommand produces python program args to launch a python file with debugpy.
func debugCommand(ctx context.Context, opts toolchain.PythonOptions) ([]string, *debugger, error) {
	err := checkForPackage(ctx, "debugpy", opts)
	if err != nil {
		var installError *NotInstalledError
		if errors.As(err, &installError) {
			return nil, nil, fmt.Errorf("debugpy is not installed. %s", installError.InstallMessage)
		}
		return nil, nil, err
	}
	logDir, err := os.MkdirTemp("", "khulnasoft-python-debugpy-")
	if err != nil {
		return nil, nil, fmt.Errorf("unable to allocate tmp dir: %w", err)
	}
	port, err := netutil.FindNextAvailablePort(preferredDebugPort)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to select a debug port: %w", err)
	}
	args := []string{}
	args = append(args, "-Xfrozen_modules=off")
	args = append(args, "-m", "debugpy", "--listen", fmt.Sprintf("127.0.0.1:%d", port))
	args = append(args, "--wait-for-client")
	args = append(args, "--log-to", logDir)
	return args, &debugger{Host: "127.0.0.1", Port: port, LogDir: logDir}, nil
}

type debugger struct {
	Host   string
	Port   int
	LogDir string
}

func (c *debugger) Cleanup() {
	contract.IgnoreError(os.RemoveAll(c.LogDir))
}

// WaitForReady waits for debugpy to be ready to accept connections.
// Returns an error if the context is canceled or the log file is unable to be tailed.
func (c *debugger) WaitForReady(ctx context.Context, pid int) error {
	logFile := filepath.Join(c.LogDir, fmt.Sprintf("debugpy.server-%d.log", pid))
	t, err := tail.TailFile(logFile, tail.Config{
		Follow: true,
		Logger: tail.DiscardingLogger,
	})
	if err != nil {
		return err
	}
	defer func() {
		contract.IgnoreError(t.Stop())
		t.Cleanup()
	}()
	ready := make(chan bool)
	go func(tailedLog *tail.Tail) {
		for line := range tailedLog.Lines {
			if line.Err != nil {
				continue
			}
			if strings.Contains(line.Text, "Adapter is accepting incoming client connections") {
				close(ready)
				break
			}
		}
	}(t)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.Dead():
		return t.Err()
	case <-ready:
		return nil
	}
}

func startDebugging(ctx context.Context, engineClient khulnasoftrpc.EngineClient, cmd *exec.Cmd, dbg *debugger) error {
	// wait for the debugger to be ready
	ctx, cancel := context.WithTimeoutCause(ctx, 1*time.Minute, errors.New("debugger startup timed out"))
	defer cancel()
	err := dbg.WaitForReady(ctx, cmd.Process.Pid)
	if err != nil {
		return err
	}

	// emit a debug configuration
	debugConfig, err := structpb.NewStruct(map[string]interface{}{
		"name":    "Khulnasoft: Program (Python)",
		"type":    "python",
		"request": "attach",
		"connect": map[string]interface{}{
			"host": dbg.Host,
			"port": dbg.Port,
		},
		"justMyCode": true,
	})
	if err != nil {
		return fmt.Errorf("failed to serialize debug configuration: %w", err)
	}
	_, err = engineClient.StartDebugging(ctx, &khulnasoftrpc.StartDebuggingRequest{
		Config:  debugConfig,
		Message: fmt.Sprintf("on port %d", dbg.Port),
	})
	if err != nil {
		return fmt.Errorf("unable to start debugging: %w", err)
	}
	return nil
}

type NotInstalledError struct {
	InstallMessage string
}

func (e *NotInstalledError) Error() string {
	return e.InstallMessage
}

func checkForPackage(ctx context.Context, pkg string, opts toolchain.PythonOptions) error {
	tc, err := toolchain.ResolveToolchain(opts)
	if err != nil {
		return err
	}
	packages, err := tc.ListPackages(ctx, true)
	if err != nil {
		return err
	}
	idx := slices.IndexFunc(packages, func(p toolchain.PythonPackage) bool { return p.Name == pkg })
	if idx < 0 {
		installCommand := fmt.Sprintf("Please install it using `poetry add %s`.", pkg)
		if opts.Toolchain != toolchain.Poetry {
			pipCommand := opts.Virtualenv + "/bin/pip install -r requirements.txt"
			if runtime.GOOS == "windows" {
				pipCommand = opts.Virtualenv + "\\Scripts\\pip install -r requirements.txt"
			}
			installCommand = fmt.Sprintf("Please add an entry for %s to requirements.txt and run `%s`", pkg, pipCommand)
		}
		//revive:disable:error-strings // This error message is user facing.
		return &NotInstalledError{
			InstallMessage: installCommand,
		}
	}
	return nil
}

// Run is RPC endpoint for LanguageRuntimeServer::Run
func (host *pythonLanguageHost) Run(ctx context.Context, req *khulnasoftrpc.RunRequest) (*khulnasoftrpc.RunResponse, error) {
	engineClient, closer, err := host.connectToEngine()
	if err != nil {
		return nil, err
	}
	defer func() {
		contract.IgnoreClose(closer)
	}()

	opts, err := parseOptions(req.Info.RootDirectory, req.Info.ProgramDirectory, req.Info.Options.AsMap())
	if err != nil {
		return nil, err
	}

	args := []string{}
	var dbg *debugger
	if req.GetAttachDebugger() {
		args, dbg, err = debugCommand(ctx, opts)
		if err != nil {
			return &khulnasoftrpc.RunResponse{
				Error: err.Error(),
			}, nil
		}
		defer dbg.Cleanup()
	}

	// Entrypoint script and arguments
	args = append(args, host.exec)
	args = append(args, host.constructArguments(req)...)

	config, err := host.constructConfig(req)
	if err != nil {
		err = fmt.Errorf("failed to serialize configuration: %w", err)
		return nil, err
	}
	configSecretKeys, err := host.constructConfigSecretKeys(req)
	if err != nil {
		err = fmt.Errorf("failed to serialize configuration secret keys: %w", err)
		return nil, err
	}

	if logging.V(5) {
		commandStr := strings.Join(args, " ")
		logging.V(5).Infoln("Language host launching process: ", host.exec, commandStr)
	}

	// Now simply spawn a process to execute the requested program, wiring up stdout/stderr directly.
	mkCmd := func(args []string) (*exec.Cmd, error) {
		tc, err := toolchain.ResolveToolchain(opts)
		if err != nil {
			return nil, err
		}

		if err := tc.ValidateVenv(ctx); err != nil {
			return nil, err
		}

		return tc.Command(ctx, args...)
	}

	cmd, err := mkCmd(args)
	if err != nil {
		return nil, err
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if config != "" || configSecretKeys != "" {
		env := cmd.Env
		if env == nil {
			env = os.Environ()
		}
		if config != "" {
			env = append(env, khulnasoftConfigVar+"="+config)
		}
		if configSecretKeys != "" {
			env = append(env, khulnasoftConfigSecretKeysVar+"="+configSecretKeys)
		}
		cmd.Env = env
	}

	// Before running the command, we might need to run typechecker first
	var typechecker string
	switch opts.Typechecker {
	case toolchain.TypeCheckerNone:
		break
	case toolchain.TypeCheckerMypy:
		typechecker = "mypy"
	case toolchain.TypeCheckerPyright:
		typechecker = "pyright"
	}

	if typechecker != "" {
		typecheckerCmd, err := mkCmd([]string{"-m", typechecker, req.Info.ProgramDirectory})
		if err != nil {
			return nil, err
		}
		typecheckerCmd.Stdout = os.Stdout
		typecheckerCmd.Stderr = os.Stderr
		err = checkForPackage(ctx, typechecker, opts)
		if err != nil {
			var installError *NotInstalledError
			if errors.As(err, &installError) {
				return nil, fmt.Errorf("The typechecker option is set to %s, but %s is not installed. %s",
					typechecker, typechecker, installError.InstallMessage)
			}
			return nil, err
		}

		if err := typecheckerCmd.Run(); err != nil {
			return nil, fmt.Errorf("%s failed: %w", typechecker, err)
		}
	}
	var errResult string
	run := func() error {
		err := cmd.Start()
		if err != nil {
			return err
		}
		if req.GetAttachDebugger() {
			// create a sub-context to cancel the startDebugging operation when the process exits.
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()
			go func() {
				err := startDebugging(ctx, engineClient, cmd, dbg)
				if err != nil {
					// kill the program if we can't start debugging.
					logging.Errorf("Unable to start debugging: %v", err)
					contract.IgnoreError(cmd.Process.Kill())
				}
			}()
		}
		return cmd.Wait()
	}
	if err := run(); err != nil {
		// Python does not explicitly flush standard out or standard error when exiting abnormally. For this reason, we
		// need to explicitly flush our output streams so that, when we exit, the engine picks up the child Python
		// process's stdout and stderr writes.
		//
		// This is especially crucial for Python because it is possible for the child Python process to crash very fast
		// if Khulnasoft is misconfigured, so we must be sure to present a high-quality error message to the user.
		contract.IgnoreError(os.Stdout.Sync())
		contract.IgnoreError(os.Stderr.Sync())
		if exiterr, ok := err.(*exec.ExitError); ok {
			// If the program ran, but exited with a non-zero error code.  This will happen often, since user
			// errors will trigger this.  So, the error message should look as nice as possible.
			if status, stok := exiterr.Sys().(syscall.WaitStatus); stok {
				switch status.ExitStatus() {
				case 0:
					// This really shouldn't happen, but if it does, we don't want to render "non-zero exit code"
					err = fmt.Errorf("program exited unexpectedly: %w", exiterr)
				case pythonProcessExitedAfterShowingUserActionableMessage:
					return &khulnasoftrpc.RunResponse{Error: "", Bail: true}, nil
				default:
					err = fmt.Errorf("program exited with non-zero exit code: %d", status.ExitStatus())
				}
			} else {
				err = fmt.Errorf("program exited unexpectedly: %w", exiterr)
			}
		} else {
			// Otherwise, we didn't even get to run the program.  This ought to never happen unless there's
			// a bug or system condition that prevented us from running the language exec.  Issue a scarier error.
			err = fmt.Errorf("problem executing program (could not run language executor): %w", err)
		}

		errResult = err.Error()
	}

	return &khulnasoftrpc.RunResponse{Error: errResult}, nil
}

// constructArguments constructs a command-line for `khulnasoft-language-python`
// by enumerating all of the optional and non-optional arguments present
// in a RunRequest.
func (host *pythonLanguageHost) constructArguments(req *khulnasoftrpc.RunRequest) []string {
	var args []string
	maybeAppendArg := func(k, v string) {
		if v != "" {
			args = append(args, "--"+k, v)
		}
	}

	maybeAppendArg("monitor", req.GetMonitorAddress())
	maybeAppendArg("engine", host.engineAddress)
	maybeAppendArg("project", req.GetProject())
	maybeAppendArg("stack", req.GetStack())
	maybeAppendArg("pwd", req.GetPwd())
	maybeAppendArg("dry_run", strconv.FormatBool(req.GetDryRun()))
	maybeAppendArg("parallel", strconv.Itoa(int(req.GetParallel())))
	maybeAppendArg("tracing", host.tracing)
	maybeAppendArg("organization", req.GetOrganization())

	// The engine should always pass a name for entry point, even if its just "." for the program directory.
	args = append(args, req.Info.EntryPoint)

	args = append(args, req.GetArgs()...)
	return args
}

// constructConfig json-serializes the configuration data given as part of a RunRequest.
func (host *pythonLanguageHost) constructConfig(req *khulnasoftrpc.RunRequest) (string, error) {
	configMap := req.GetConfig()
	if configMap == nil {
		return "", nil
	}

	configJSON, err := json.Marshal(configMap)
	if err != nil {
		return "", err
	}

	return string(configJSON), nil
}

// constructConfigSecretKeys JSON-serializes the list of keys that contain secret values given as part of
// a RunRequest.
func (host *pythonLanguageHost) constructConfigSecretKeys(req *khulnasoftrpc.RunRequest) (string, error) {
	configSecretKeys := req.GetConfigSecretKeys()
	if configSecretKeys == nil {
		return "[]", nil
	}

	configSecretKeysJSON, err := json.Marshal(configSecretKeys)
	if err != nil {
		return "", err
	}

	return string(configSecretKeysJSON), nil
}

func (host *pythonLanguageHost) GetPluginInfo(ctx context.Context, req *emptypb.Empty) (*khulnasoftrpc.PluginInfo, error) {
	return &khulnasoftrpc.PluginInfo{
		Version: version.Version,
	}, nil
}

// validateVersion checks that python is running a valid version. If a version
// is invalid, it prints to os.Stderr. This is interpreted as diagnostic message
// by the Khulnasoft CLI program.
func validateVersion(ctx context.Context, options toolchain.PythonOptions) {
	var versionCmd *exec.Cmd
	var err error
	versionArgs := []string{"--version"}

	tc, err := toolchain.ResolveToolchain(options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to configure python toolchain: %s\n", err)
		return
	}

	versionCmd, err = tc.Command(ctx, versionArgs...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create python version command: %s\n", err)
		return
	}
	var out []byte
	if out, err = versionCmd.Output(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to resolve python version command: %s\n", err)
		return
	}
	version := strings.TrimSpace(strings.TrimPrefix(string(out), "Python "))
	version = removeReleaseCandidateSuffix(version)
	parsed, err := semver.Parse(version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse python version: '%s'\n", version)
		return
	}
	if parsed.LT(minimumSupportedPythonVersion) {
		fmt.Fprintf(os.Stderr, "Khulnasoft does not support Python %s."+
			" Please upgrade to at least %s\n", parsed, minimumSupportedPythonVersion)
	} else if parsed.LT(eolPythonVersion) {
		fmt.Fprintf(os.Stderr, "Python %d.%d is approaching EOL and will not be supported in Khulnasoft soon."+
			" Check %s for more details\n", parsed.Major,
			parsed.Minor, eolPythonVersionIssue)
	}
}

func (host *pythonLanguageHost) InstallDependencies(
	req *khulnasoftrpc.InstallDependenciesRequest, server khulnasoftrpc.LanguageRuntime_InstallDependenciesServer,
) error {
	opts, err := parseOptions(req.Info.RootDirectory, req.Info.ProgramDirectory, req.Info.Options.AsMap())
	if err != nil {
		return err
	}

	closer, stdout, stderr, err := rpcutil.MakeInstallDependenciesStreams(server, req.IsTerminal)
	if err != nil {
		return err
	}
	// best effort close, but we try an explicit close and error check at the end as well
	defer closer.Close()

	stdout.Write([]byte("Installing dependencies...\n\n"))

	tc, err := toolchain.ResolveToolchain(opts)
	if err != nil {
		return err
	}
	if err := tc.InstallDependencies(server.Context(), req.Info.ProgramDirectory, req.UseLanguageVersionTools,
		true /*showOutput*/, stdout, stderr); err != nil {
		return err
	}

	stdout.Write([]byte("Finished installing dependencies\n\n"))

	return closer.Close()
}

func (host *pythonLanguageHost) RuntimeOptionsPrompts(ctx context.Context,
	req *khulnasoftrpc.RuntimeOptionsRequest,
) (*khulnasoftrpc.RuntimeOptionsResponse, error) {
	rawOpts := req.Info.Options.AsMap()

	var prompts []*khulnasoftrpc.RuntimeOptionPrompt

	toolchain, hasToolchain := rawOpts["toolchain"]

	if !hasToolchain {
		pipOption := &khulnasoftrpc.RuntimeOptionPrompt_RuntimeOptionValue{
			PromptType:  khulnasoftrpc.RuntimeOptionPrompt_STRING,
			StringValue: "pip",
			DisplayName: "pip",
		}
		// Pip is always available in a Python installation or virtual environment.
		choices := []*khulnasoftrpc.RuntimeOptionPrompt_RuntimeOptionValue{pipOption}
		choices = append(choices, plugin.MakeExecutablePromptChoices("poetry", "uv")...)
		prompts = append(prompts, &khulnasoftrpc.RuntimeOptionPrompt{
			Key:         "toolchain",
			Description: "The toolchain to use for installing dependencies and running the program",
			PromptType:  khulnasoftrpc.RuntimeOptionPrompt_STRING,
			Choices:     choices,
			Default:     pipOption,
		})
	}

	if hasToolchain && toolchain == "pip" {
		// If we are using the pip toolchain, set virtualenv to venv by default.
		if _, hasVenv := rawOpts["virtualenv"]; !hasVenv {
			prompts = append(prompts, &khulnasoftrpc.RuntimeOptionPrompt{
				Key:         "virtualenv",
				Description: "The virtualenv to use",
				PromptType:  khulnasoftrpc.RuntimeOptionPrompt_STRING,
				Choices: []*khulnasoftrpc.RuntimeOptionPrompt_RuntimeOptionValue{
					{StringValue: "venv", PromptType: khulnasoftrpc.RuntimeOptionPrompt_STRING},
				},
				Default: &khulnasoftrpc.RuntimeOptionPrompt_RuntimeOptionValue{
					PromptType:  khulnasoftrpc.RuntimeOptionPrompt_STRING,
					StringValue: "venv",
				},
			})
		}
	}

	return &khulnasoftrpc.RuntimeOptionsResponse{
		Prompts: prompts,
	}, nil
}

func (host *pythonLanguageHost) About(ctx context.Context,
	req *khulnasoftrpc.AboutRequest,
) (*khulnasoftrpc.AboutResponse, error) {
	// Previously we did not pass any arguments to About and we always used the default python command.
	opts := toolchain.PythonOptions{Toolchain: toolchain.Pip}
	if req != nil && req.Info != nil {
		aboutOpts, err := parseOptions(req.Info.RootDirectory, req.Info.ProgramDirectory, req.Info.Options.AsMap())
		if err != nil {
			return nil, err
		}
		opts = aboutOpts
	}

	tc, err := toolchain.ResolveToolchain(opts)
	if err != nil {
		return nil, err
	}

	info, err := tc.About(ctx)
	if err != nil {
		return nil, err
	}

	return &khulnasoftrpc.AboutResponse{
		Executable: info.Executable,
		Version:    info.Version,
	}, nil
}

func (host *pythonLanguageHost) GetProgramDependencies(
	ctx context.Context, req *khulnasoftrpc.GetProgramDependenciesRequest,
) (*khulnasoftrpc.GetProgramDependenciesResponse, error) {
	opts, err := parseOptions(req.Info.RootDirectory, req.Info.ProgramDirectory, req.Info.Options.AsMap())
	if err != nil {
		return nil, err
	}

	tc, err := toolchain.ResolveToolchain(opts)
	if err != nil {
		return nil, err
	}
	result, err := tc.ListPackages(ctx, req.TransitiveDependencies /* transitive */)
	if err != nil {
		return nil, fmt.Errorf("failed to get python dependencies: %w", err)
	}

	dependencies := make([]*khulnasoftrpc.DependencyInfo, len(result))
	for i, dep := range result {
		dependencies[i] = &khulnasoftrpc.DependencyInfo{
			Name:    dep.Name,
			Version: dep.Version,
		}
	}

	return &khulnasoftrpc.GetProgramDependenciesResponse{
		Dependencies: dependencies,
	}, nil
}

func (host *pythonLanguageHost) RunPlugin(
	req *khulnasoftrpc.RunPluginRequest, server khulnasoftrpc.LanguageRuntime_RunPluginServer,
) error {
	logging.V(5).Infof("Attempting to run python plugin in %s", req.Info.ProgramDirectory)

	opts, err := parseOptions(req.Info.RootDirectory, req.Info.ProgramDirectory, req.Info.Options.AsMap())
	if err != nil {
		return err
	}

	args := []string{req.Info.ProgramDirectory}
	args = append(args, req.Args...)

	tc, err := toolchain.ResolveToolchain(opts)
	if err != nil {
		return err
	}
	cmd, err := tc.Command(server.Context(), args...)
	if err != nil {
		return err
	}

	closer, stdout, stderr, err := rpcutil.MakeRunPluginStreams(server, false)
	if err != nil {
		return err
	}
	// best effort close, but we try an explicit close and error check at the end as well
	defer closer.Close()

	cmd.Dir = req.Pwd
	cmd.Env = append(cmd.Env, req.Env...)
	cmd.Stdout, cmd.Stderr = stdout, stderr

	if err = cmd.Run(); err != nil {
		var exiterr *exec.ExitError
		if errors.As(err, &exiterr) {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return server.Send(&khulnasoftrpc.RunPluginResponse{
					//nolint:gosec // WaitStatus always uses the lower 8 bits for the exit code.
					Output: &khulnasoftrpc.RunPluginResponse_Exitcode{Exitcode: int32(status.ExitStatus())},
				})
			}
			return fmt.Errorf("program exited unexpectedly: %w", exiterr)
		}
		return fmt.Errorf("problem executing plugin program (could not run language executor): %w", err)
	}

	return closer.Close()
}

func (host *pythonLanguageHost) GenerateProject(
	ctx context.Context, req *khulnasoftrpc.GenerateProjectRequest,
) (*khulnasoftrpc.GenerateProjectResponse, error) {
	loader, err := schema.NewLoaderClient(req.LoaderTarget)
	if err != nil {
		return nil, err
	}
	defer loader.Close()

	var extraOptions []pcl.BindOption
	if !req.Strict {
		extraOptions = append(extraOptions, pcl.NonStrictBindOptions()...)
	}

	// for python, prefer output-versioned invokes
	extraOptions = append(extraOptions, pcl.PreferOutputVersionedInvokes)

	program, diags, err := pcl.BindDirectory(req.SourceDirectory, schema.NewCachedLoader(loader), extraOptions...)
	if err != nil {
		return nil, err
	}
	if diags.HasErrors() {
		rpcDiagnostics := plugin.HclDiagnosticsToRPCDiagnostics(diags)

		return &khulnasoftrpc.GenerateProjectResponse{
			Diagnostics: rpcDiagnostics,
		}, nil
	}

	var project workspace.Project
	if err := json.Unmarshal([]byte(req.Project), &project); err != nil {
		return nil, err
	}

	err = codegen.GenerateProject(req.TargetDirectory, project, program, req.LocalDependencies, host.typechecker)
	if err != nil {
		return nil, err
	}

	rpcDiagnostics := plugin.HclDiagnosticsToRPCDiagnostics(diags)

	return &khulnasoftrpc.GenerateProjectResponse{
		Diagnostics: rpcDiagnostics,
	}, nil
}

func (host *pythonLanguageHost) GenerateProgram(
	ctx context.Context, req *khulnasoftrpc.GenerateProgramRequest,
) (*khulnasoftrpc.GenerateProgramResponse, error) {
	loader, err := schema.NewLoaderClient(req.LoaderTarget)
	if err != nil {
		return nil, err
	}
	defer loader.Close()

	parser := hclsyntax.NewParser()
	// Load all .pp files in the directory
	for path, contents := range req.Source {
		err = parser.ParseFile(strings.NewReader(contents), path)
		if err != nil {
			return nil, err
		}
		diags := parser.Diagnostics
		if diags.HasErrors() {
			return nil, diags
		}
	}

	bindOptions := []pcl.BindOption{
		pcl.Loader(schema.NewCachedLoader(loader)),
		// for python, prefer output-versioned invokes
		pcl.PreferOutputVersionedInvokes,
	}

	if !req.Strict {
		bindOptions = append(bindOptions, pcl.NonStrictBindOptions()...)
	}

	program, diags, err := pcl.BindProgram(parser.Files, bindOptions...)
	if err != nil {
		return nil, err
	}

	rpcDiagnostics := plugin.HclDiagnosticsToRPCDiagnostics(diags)
	if diags.HasErrors() {
		return &khulnasoftrpc.GenerateProgramResponse{
			Diagnostics: rpcDiagnostics,
		}, nil
	}
	if program == nil {
		return nil, errors.New("internal error program was nil")
	}

	files, diags, err := codegen.GenerateProgram(program)
	if err != nil {
		return nil, err
	}
	rpcDiagnostics = append(rpcDiagnostics, plugin.HclDiagnosticsToRPCDiagnostics(diags)...)

	return &khulnasoftrpc.GenerateProgramResponse{
		Source:      files,
		Diagnostics: rpcDiagnostics,
	}, nil
}

func (host *pythonLanguageHost) GeneratePackage(
	ctx context.Context, req *khulnasoftrpc.GeneratePackageRequest,
) (*khulnasoftrpc.GeneratePackageResponse, error) {
	loader, err := schema.NewLoaderClient(req.LoaderTarget)
	if err != nil {
		return nil, err
	}
	defer loader.Close()

	var spec schema.PackageSpec
	err = json.Unmarshal([]byte(req.Schema), &spec)
	if err != nil {
		return nil, err
	}

	pkg, diags, err := schema.BindSpec(spec, schema.NewCachedLoader(loader))
	if err != nil {
		return nil, err
	}
	rpcDiagnostics := plugin.HclDiagnosticsToRPCDiagnostics(diags)
	if diags.HasErrors() {
		return &khulnasoftrpc.GeneratePackageResponse{
			Diagnostics: rpcDiagnostics,
		}, nil
	}

	files, err := codegen.GeneratePackage("khulnasoft-language-python", pkg, req.ExtraFiles)
	if err != nil {
		return nil, err
	}

	for filename, data := range files {
		outPath := filepath.Join(req.Directory, filename)

		err := os.MkdirAll(filepath.Dir(outPath), 0o700)
		if err != nil {
			return nil, fmt.Errorf("could not create output directory %s: %w", filepath.Dir(filename), err)
		}

		err = os.WriteFile(outPath, data, 0o600)
		if err != nil {
			return nil, fmt.Errorf("could not write output file %s: %w", filename, err)
		}
	}

	return &khulnasoftrpc.GeneratePackageResponse{
		Diagnostics: rpcDiagnostics,
	}, nil
}

// removeReleaseCandidateSuffix removes any "rc" suffix from a semantic version string.
func removeReleaseCandidateSuffix(version string) string {
	re := regexp.MustCompile(`-?rc\d+$`)
	return re.ReplaceAllString(version, "")
}
