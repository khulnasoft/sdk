// Copyright 2016-2022, KhulnaSoft, Ltd.
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

package workspace

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/khulnasoft/sdk/go/common/encoding"
	"github.com/khulnasoft/sdk/go/common/tokens"
	"github.com/khulnasoft/sdk/go/common/util/fsutil"
)

const (
	// BackupDir is the name of the folder where backup stack information is stored.
	BackupDir = "backups"
	// BookkeepingDir is the name of our bookkeeping folder, we store state here (like .git for git).
	BookkeepingDir = ".khulnasoft"
	// ConfigDir is the name of the folder that holds local configuration information.
	ConfigDir = "config"
	// GitDir is the name of the folder git uses to store information.
	GitDir = ".git"
	// HistoryDir is the name of the directory that holds historical information for projects.
	HistoryDir = "history"
	// PluginDir is the name of the directory containing plugins.
	PluginDir = "plugins"
	// PolicyDir is the name of the directory that holds policy packs.
	PolicyDir = "policies"
	// StackDir is the name of the directory that holds stack information for projects.
	StackDir = "stacks"
	// LockDir is the name of the directory that holds locking information for projects.
	LockDir = "locks"
	// TemplateDir is the name of the directory containing templates.
	TemplateDir = "templates"
	// TemplatePolicyDir is the name of the directory containing templates for Policy Packs.
	TemplatePolicyDir = "templates-policy"
	// WorkspaceDir is the name of the directory that holds workspace information for projects.
	WorkspaceDir = "workspaces"

	// IgnoreFile is the name of the file that we use to control what to upload to the service.
	IgnoreFile = ".khulnasoftignore"

	// ProjectFile is the base name of a project file.
	ProjectFile = "Khulnasoft"
	// DeploymentSuffix is the base suffix for deployment settings files (e.g. "Khulnasoft.<stack>.deploy.yaml").
	DeploymentSuffix = "deploy"
	// RepoFile is the name of the file that holds information specific to the entire repository.
	RepoFile = "settings.json"
	// WorkspaceFile is the name of the file that holds workspace information.
	WorkspaceFile = "workspace.json"
	// CachedVersionFile is the name of the file we use to store when we last checked if the CLI was out of date
	CachedVersionFile = ".cachedVersionInfo"

	// KhulnasoftHomeEnvVar is a path to the '.khulnasoft' folder with plugins, access token, etc.
	// The folder can have any name, not necessarily '.khulnasoft'.
	// It defaults to the '<user's home>/.khulnasoft' if not specified.
	KhulnasoftHomeEnvVar = "KHULNASOFT_HOME"

	// PolicyPackFile is the base name of a Khulnasoft policy pack file.
	PolicyPackFile = "KhulnasoftPolicy"
)

// DetectProjectPath locates the closest project from the current working directory, or an error if not found.
func DetectProjectPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path, err := DetectProjectPathFrom(dir)
	if err != nil {
		return "", err
	}

	return path, nil
}

// DetectProjectStackPath returns the name of the file to store stack specific project settings in. We place stack
// specific settings next to the Khulnasoft.yaml file, named like: Khulnasoft.<stack-name>.yaml
func DetectProjectStackPath(stackName tokens.QName) (*Project, string, error) {
	proj, projPath, err := DetectProjectAndPath()
	if err != nil {
		return nil, "", err
	}

	fileName := fmt.Sprintf("%s.%s%s", ProjectFile, qnameFileName(stackName), filepath.Ext(projPath))

	if proj.StackConfigDir != "" {
		return proj, filepath.Join(filepath.Dir(projPath), proj.StackConfigDir, fileName), nil
	}

	return proj, filepath.Join(filepath.Dir(projPath), fileName), nil
}

func DetectProjectStackDeploymentPath(stackName tokens.QName) (string, error) {
	proj, projPath, err := DetectProjectAndPath()
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%s.%s.%s%s", ProjectFile, qnameFileName(stackName), DeploymentSuffix, filepath.Ext(projPath))

	if proj.StackConfigDir != "" {
		return filepath.Join(filepath.Dir(projPath), proj.StackConfigDir, fileName), nil
	}

	return filepath.Join(filepath.Dir(projPath), fileName), nil
}

var ErrProjectNotFound = errors.New("no project file found")

// DetectProjectPathFrom locates the closest project from the given path, searching "upwards" in the directory
// hierarchy.  If no project is found, an empty path is returned.
func DetectProjectPathFrom(dir string) (string, error) {
	path, err := fsutil.WalkUp(dir, isProject, func(s string) bool {
		return true
	})
	// We special case permission errors to cause ErrProjectNotFound to return from this function. This is so
	// users can run khulnasoft with unreadable root directories.
	if errors.Is(err, fs.ErrPermission) {
		err = nil
	}

	if err != nil {
		return "", fmt.Errorf("failed to locate Khulnasoft.yaml project file: %w", err)
	}

	if path == "" {
		// Embed/wrap ErrProjectNotFound
		return "", fmt.Errorf(
			"no Khulnasoft.yaml project file found (searching upwards from %s). If you have not "+
				"created a project yet, use `khulnasoft new` to do so: %w", dir, ErrProjectNotFound)
	}
	return path, nil
}

// DetectPolicyPackPathFrom locates the closest Khulnasoft policy project from the given path,
// searching "upwards" in the directory hierarchy.  If no project is found, an empty path is
// returned.
func DetectPolicyPackPathFrom(path string) (string, error) {
	return fsutil.WalkUp(path, isPolicyPack, func(s string) bool {
		return true
	})
}

// DetectPolicyPackPathAt locates the KhulnasoftPolicy file at the given path. If no project is found, an empty path is
// returned. Unlike DetectPolicyPackPathFrom, this function does not search upwards in the directory hierarchy.
func DetectPolicyPackPathAt(path string) (string, error) {
	for _, ext := range encoding.Exts {
		packPath := filepath.Join(path, PolicyPackFile+ext)
		info, err := os.Stat(packPath)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				continue
			}
			return "", err
		}
		if info.IsDir() {
			return "", nil
		}
		return packPath, nil
	}
	return "", nil
}

// DetectProject loads the closest project from the current working directory, or an error if not found.
func DetectProject() (*Project, error) {
	proj, _, err := DetectProjectAndPath()
	return proj, err
}

func DetectProjectStack(stackName tokens.QName) (*ProjectStack, error) {
	project, path, err := DetectProjectStackPath(stackName)
	if err != nil {
		return nil, err
	}

	return LoadProjectStack(project, path)
}

func DetectProjectStackDeployment(stackName tokens.QName) (*ProjectStackDeployment, error) {
	path, err := DetectProjectStackDeploymentPath(stackName)
	if err != nil {
		return nil, err
	}

	return LoadProjectStackDeployment(path)
}

// DetectProjectAndPath loads the closest package from the current working directory, or an error if not found.  It
// also returns the path where the package was found.
func DetectProjectAndPath() (*Project, string, error) {
	path, err := DetectProjectPath()
	if err != nil {
		return nil, "", err
	} else if path == "" {
		return nil, "", errors.New("no Khulnasoft project found in the current working directory. " +
			"Move to a directory with a Khulnasoft project or try creating a project first with `khulnasoft new`.")
	}

	proj, err := LoadProject(path)
	return proj, path, err
}

// SaveProject saves the project file on top of the existing one, using the standard location.
func SaveProject(proj *Project) error {
	path, err := DetectProjectPath()
	if err != nil {
		return err
	}

	return proj.Save(path)
}

func SaveProjectStack(stackName tokens.QName, stack *ProjectStack) error {
	_, path, err := DetectProjectStackPath(stackName)
	if err != nil {
		return err
	}

	return stack.Save(path)
}

func SaveProjectStackDeployment(stackName tokens.QName, deployment *ProjectStackDeployment) error {
	path, err := DetectProjectStackDeploymentPath(stackName)
	if err != nil {
		return err
	}

	return deployment.Save(path)
}

// isProject returns true if the path references what appears to be a valid project.  If problems are detected -- like
// an incorrect extension -- they are logged to the provided diag.Sink (if non-nil).
func isProject(path string) bool {
	return isMarkupFile(path, ProjectFile)
}

// isPolicyPack returns true if the path references what appears to be a valid policy pack project.
// If problems are detected -- like an incorrect extension -- they are logged to the provided
// diag.Sink (if non-nil).
func isPolicyPack(path string) bool {
	return isMarkupFile(path, PolicyPackFile)
}

func isMarkupFile(path string, expect string) bool {
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		// Missing files and directories can't be markup files.
		return false
	}

	// Ensure the base name is expected.
	name := info.Name()
	ext := filepath.Ext(name)
	base := strings.TrimSuffix(name, ext)
	if base != expect {
		return false
	}

	// Check all supported extensions.
	for _, mext := range encoding.Exts {
		if name == expect+mext {
			return true
		}
	}

	return false
}

// GetCachedVersionFilePath returns the location where the CLI caches information from khulnasoft.com on the newest
// available version of the CLI
func GetCachedVersionFilePath() (string, error) {
	return GetKhulnasoftPath(CachedVersionFile)
}

// GetKhulnasoftHomeDir returns the path of the '.khulnasoft' folder where Khulnasoft puts its artifacts.
func GetKhulnasoftHomeDir() (string, error) {
	// Allow the folder we use to be overridden by an environment variable
	dir := os.Getenv(KhulnasoftHomeEnvVar)
	if dir != "" {
		return dir, nil
	}

	// Otherwise, use the current user's home dir + .khulnasoft
	user, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("getting current user: %w", err)
	}

	if user == nil || user.HomeDir == "" {
		return "", fmt.Errorf("could not find user home directory, set %s", KhulnasoftHomeEnvVar)
	}

	return filepath.Join(user.HomeDir, BookkeepingDir), nil
}

// GetKhulnasoftPath returns the path to a file or directory under the '.khulnasoft' folder. It joins the path of
// the '.khulnasoft' folder with elements passed as arguments.
func GetKhulnasoftPath(elem ...string) (string, error) {
	homeDir, err := GetKhulnasoftHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(append([]string{homeDir}, elem...)...), nil
}
