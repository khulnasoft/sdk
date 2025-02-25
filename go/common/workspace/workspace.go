// Copyright 2016-2021, KhulnaSoft, Ltd.
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
	//nolint:gosec
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/khulnasoft/sdk/go/common/tokens"
	"github.com/khulnasoft/sdk/go/common/util/contract"
)

// W offers functionality for interacting with Khulnasoft workspaces.
type W interface {
	Settings() *Settings // returns a mutable pointer to the optional workspace settings info.
	Save() error         // saves any modifications to the workspace.
}

type projectWorkspace struct {
	name     tokens.PackageName // the package this workspace is associated with.
	project  string             // the path to the Khulnasoft.[yaml|json] file for this project.
	settings *Settings          // settings for this workspace.
}

var (
	cache      = make(map[string]W)
	cacheMutex sync.RWMutex
)

func loadFromCache(key string) (W, bool) {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()

	w, ok := cache[key]
	return w, ok
}

func upsertIntoCache(key string, w W) {
	contract.Requiref(w != nil, "w", "cannot be nil")

	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	cache[key] = w
}

// New creates a new workspace using the current working directory.
func New() (W, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return NewFrom(cwd)
}

// NewFrom creates a new Khulnasoft workspace in the given directory. Requires a Khulnasoft.yaml file be present in the
// folder hierarchy between dir and the .khulnasoft folder.
func NewFrom(dir string) (W, error) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	dir = absDir

	if w, ok := loadFromCache(dir); ok {
		return w, nil
	}

	path, err := DetectProjectPathFrom(dir)
	if err != nil {
		return nil, err
	} else if path == "" {
		return nil, fmt.Errorf("no Khulnasoft.yaml project file found (searching upwards from %s). If you have not "+
			"created a project yet, use `khulnasoft new` to do so", dir)
	}

	proj, err := LoadProject(path)
	if err != nil {
		return nil, err
	}

	w := &projectWorkspace{
		name:    proj.Name,
		project: path,
	}

	err = w.readSettings()
	if err != nil {
		return nil, fmt.Errorf("unable to read workspace settings: %w", err)
	}

	upsertIntoCache(dir, w)
	return w, nil
}

func (pw *projectWorkspace) Settings() *Settings {
	return pw.settings
}

func (pw *projectWorkspace) Save() error {
	settingsFile := pw.settingsPath()

	// If the settings file is empty, don't write an new one, and delete the old one if present. Since we put workspaces
	// under ~/.khulnasoft/workspaces, cleaning them out when possible prevents us from littering a bunch of files in the
	// home directory.
	if pw.settings.IsEmpty() {
		err := os.Remove(settingsFile)
		if err != nil && !os.IsNotExist(err) {
			return err
		}
		return nil
	}

	err := os.MkdirAll(filepath.Dir(settingsFile), 0o700)
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(pw.settings, "", "    ")
	if err != nil {
		return err
	}
	return atomicWriteFile(settingsFile, b)
}

// atomicWriteFile provides a rename based atomic write through a temporary file.
func atomicWriteFile(path string, b []byte) error {
	tmp, err := os.CreateTemp(filepath.Dir(path), filepath.Base(path))
	if err != nil {
		return fmt.Errorf("failed to create temporary file %s: %w", path, err)
	}
	defer func() { contract.Ignore(os.Remove(tmp.Name())) }()

	if err = tmp.Chmod(0o600); err != nil {
		return fmt.Errorf("failed to set temporary file permission: %w", err)
	}
	if _, err = tmp.Write(b); err != nil {
		return fmt.Errorf("failed to write to temporary file: %w", err)
	}
	if err = tmp.Sync(); err != nil {
		return err
	}
	if err = tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), path)
}

func (pw *projectWorkspace) readSettings() error {
	settingsPath := pw.settingsPath()

	b, err := os.ReadFile(settingsPath)
	if err != nil && os.IsNotExist(err) {
		// not an error to not have an existing settings file.
		pw.settings = &Settings{}
		return nil
	} else if err != nil {
		return err
	}

	var settings Settings

	err = json.Unmarshal(b, &settings)
	if err != nil {
		return fmt.Errorf("could not parse file %s: %w", settingsPath, err)
	}

	pw.settings = &settings
	return nil
}

func (pw *projectWorkspace) settingsPath() string {
	uniqueFileName := string(pw.name) + "-" + sha1HexString(pw.project) + "-" + WorkspaceFile
	path, err := GetKhulnasoftPath(WorkspaceDir, uniqueFileName)
	contract.AssertNoErrorf(err, "could not get workspace path")
	return path
}

// sha1HexString returns a hex string of the sha1 hash of value.
func sha1HexString(value string) string {
	//nolint:gosec
	h := sha1.New()
	_, err := h.Write([]byte(value))
	contract.AssertNoErrorf(err, "error hashing string")
	return hex.EncodeToString(h.Sum(nil))
}

// qnameFileName takes a qname and cleans it for use as a filename (by replacing tokens.QNameDelimter with a dash)
func qnameFileName(nm tokens.QName) string {
	return strings.ReplaceAll(string(nm), tokens.QNameDelimiter, "-")
}
