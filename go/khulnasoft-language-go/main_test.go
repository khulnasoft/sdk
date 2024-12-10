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

package main

import (
	"context"
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/khulnasoft/sdk/go/common/resource/plugin"
	"github.com/khulnasoft/sdk/go/common/testing/iotest"
	"github.com/khulnasoft/sdk/go/common/util/fsutil"
	khulnasoftrpc "github.com/khulnasoft/sdk/proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseRunParams(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc    string
		give    []string
		want    runParams
		wantErr string // non-empty if we expect an error
	}{
		{
			desc:    "no arguments",
			wantErr: "missing required engine RPC address argument",
		},
		{
			desc: "no options",
			give: []string{"localhost:1234"},
			want: runParams{
				engineAddress: "localhost:1234",
			},
		},
		{
			desc: "tracing",
			give: []string{"-tracing", "foo.trace", "localhost:1234"},
			want: runParams{
				tracing:       "foo.trace",
				engineAddress: "localhost:1234",
			},
		},
		{
			desc: "binary",
			give: []string{"-binary", "foo", "localhost:1234"},
			want: runParams{
				engineAddress: "localhost:1234",
			},
		},
		{
			desc: "buildTarget",
			give: []string{"-buildTarget", "foo", "localhost:1234"},
			want: runParams{
				engineAddress: "localhost:1234",
			},
		},
		{
			desc: "root",
			give: []string{"-root", "path/to/root", "localhost:1234"},
			want: runParams{
				engineAddress: "localhost:1234",
			},
		},
		{
			desc:    "unknown option",
			give:    []string{"-unknown-option", "bar", "localhost:1234"},
			wantErr: "flag provided but not defined: -unknown-option",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			// Use a FlagSet with ContinueOnError for each case
			// instead of using the global flag set.
			//
			// The global flag set uses flag.ExitOnError,
			// so it cannot validate error cases during tests.
			fset := flag.NewFlagSet(t.Name(), flag.ContinueOnError)
			fset.SetOutput(iotest.LogWriter(t))

			got, err := parseRunParams(fset, tt.give)
			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &tt.want, got)
			}
		})
	}
}

func TestGetPackage(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name          string
		Mod           *modInfo
		Expected      *khulnasoftrpc.PackageDependency
		ExpectedError string
		JSON          *plugin.KhulnasoftPluginJSON
		JSONPath      string
	}{
		{
			Name: "valid-khulnasoft-mod",
			Mod: &modInfo{
				Path:    "github.com/khulnasoft/khulnasoft-aws/sdk",
				Version: "v1.29.0",
			},
			Expected: &khulnasoftrpc.PackageDependency{
				Name:    "aws",
				Version: "v1.29.0",
			},
		},
		{
			Name: "khulnasoft-pseduo-version-plugin",
			Mod: &modInfo{
				Path:    "github.com/khulnasoft/khulnasoft-aws/sdk",
				Version: "v1.29.1-0.20200403140640-efb5e2a48a86",
			},
			Expected: &khulnasoftrpc.PackageDependency{
				Name:    "aws",
				Version: "v1.29.0",
			},
		},
		{
			Name: "non-khulnasoft-mod",
			Mod: &modInfo{
				Path:    "github.com/moolumi/khulnasoft-aws/sdk",
				Version: "v1.29.0",
			},
			ExpectedError: "module is not a khulnasoft provider",
		},
		{
			Name: "invalid-version-module",
			Mod: &modInfo{
				Path:    "github.com/khulnasoft/khulnasoft-aws/sdk",
				Version: "42-42-42",
			},
			ExpectedError: "module does not have semver compatible version",
		},
		{
			Name: "khulnasoft-khulnasoft-mod",
			Mod: &modInfo{
				Path:    "github.com/khulnasoft/sdk",
				Version: "v1.14.0",
			},
			ExpectedError: "module is not a khulnasoft provider",
		},
		{
			Name: "beta-khulnasoft-module",
			Mod: &modInfo{
				Path:    "github.com/khulnasoft/khulnasoft-aws/sdk",
				Version: "v2.0.0-beta.1",
			},
			Expected: &khulnasoftrpc.PackageDependency{
				Name:    "aws",
				Version: "v2.0.0-beta.1",
			},
		},
		{
			Name: "non-zero-patch-module", Mod: &modInfo{
				Path:    "github.com/khulnasoft/khulnasoft-kubernetes/sdk",
				Version: "v1.5.8",
			},
			Expected: &khulnasoftrpc.PackageDependency{
				Name:    "kubernetes",
				Version: "v1.5.8",
			},
		},
		{
			Name: "khulnasoftplugin",
			Mod: &modInfo{
				Path:    "github.com/me/myself/i",
				Version: "invalid-Version",
			},
			Expected: &khulnasoftrpc.PackageDependency{
				Name:    "thing1",
				Version: "v1.2.3",
				Server:  "myserver.com",
			},
			JSON: &plugin.KhulnasoftPluginJSON{
				Resource: true,
				Name:     "thing1",
				Version:  "v1.2.3",
				Server:   "myserver.com",
			},
		},
		{
			Name:          "non-resource",
			Mod:           &modInfo{},
			ExpectedError: "module is not a khulnasoft provider",
			JSON: &plugin.KhulnasoftPluginJSON{
				Resource: false,
			},
		},
		{
			Name: "missing-khulnasoftplugin",
			Mod: &modInfo{
				Dir: "/not/real",
			},
			ExpectedError: "module is not a khulnasoft provider",
			JSON: &plugin.KhulnasoftPluginJSON{
				Name:    "thing2",
				Version: "v1.2.3",
			},
		},
		{
			Name: "khulnasoftplugin-go-lookup",
			Mod: &modInfo{
				Path:    "github.com/me/myself",
				Version: "v1.2.3",
			},
			JSON: &plugin.KhulnasoftPluginJSON{
				Name:     "name",
				Resource: true,
			},
			JSONPath: "go",
			Expected: &khulnasoftrpc.PackageDependency{
				Name:    "name",
				Version: "v1.2.3",
			},
		},
		{
			Name: "khulnasoftplugin-go-name-lookup",
			Mod: &modInfo{
				Path:    "github.com/me/myself",
				Version: "v1.2.3",
			},
			JSON: &plugin.KhulnasoftPluginJSON{
				Name:     "name",
				Resource: true,
			},
			JSONPath: filepath.Join("go", "name"),
			Expected: &khulnasoftrpc.PackageDependency{
				Name:    "name",
				Version: "v1.2.3",
			},
		},
		{
			Name: "khulnasoftplugin-nested-too-deep",
			Mod: &modInfo{
				Path:    "path.com/here",
				Version: "v0.0",
			},
			JSONPath: filepath.Join("go", "valid", "invalid"),
			JSON: &plugin.KhulnasoftPluginJSON{
				Name:     "name",
				Resource: true,
			},
			ExpectedError: "module is not a khulnasoft provider",
		},
		{
			Name: "nested-wrong-folder",
			Mod: &modInfo{
				Path:    "path.com/here",
				Version: "v0.0",
			},
			JSONPath: filepath.Join("invalid", "valid"),
			JSON: &plugin.KhulnasoftPluginJSON{
				Name:     "name",
				Resource: true,
			},
			ExpectedError: "module is not a khulnasoft provider",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			cwd := t.TempDir()
			if c.Mod.Dir == "" {
				c.Mod.Dir = cwd
			}
			if c.JSON != nil {
				path := filepath.Join(cwd, c.JSONPath)
				err := os.MkdirAll(path, 0o700)
				assert.NoErrorf(t, err, "Failed to setup test folder %s", path)
				bytes, err := c.JSON.JSON()
				assert.NoError(t, err, "Failed to setup test khulnasoft-plugin.json")
				err = os.WriteFile(filepath.Join(path, "khulnasoft-plugin.json"), bytes, 0o600)
				assert.NoError(t, err, "Failed to write khulnasoft-plugin.json")
			}

			actual, err := c.Mod.getPackage(t.TempDir())
			if c.ExpectedError != "" {
				assert.EqualError(t, err, c.ExpectedError)
			} else {
				// Kind must be resource. We can thus exclude it from the test.
				if c.Expected.Kind == "" {
					c.Expected.Kind = "resource"
				}
				assert.NoError(t, err)
				assert.Equal(t, c.Expected, actual)
			}
		})
	}
}

func TestPluginsAndDependencies_moduleMode(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	require.NoError(t,
		fsutil.CopyFile(root, filepath.Join("testdata", "sample"), nil),
		"copy test data")

	testPluginsAndDependencies(t, filepath.Join(root, "prog"))
}

// Test for https://github.com/khulnasoft/khulnasoft/issues/12526.
// Validates that if a Khulnasoft program has vendored its dependencies,
// the language host can still find the plugin and run the program.
func TestPluginsAndDependencies_vendored(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	require.NoError(t,
		fsutil.CopyFile(root, filepath.Join("testdata", "sample"), nil),
		"copy test data")

	progDir := filepath.Join(root, "prog")

	// Vendor the dependencies and nuke the sources
	// to ensure that the language host can only use the vendored version.
	cmd := exec.Command("go", "mod", "vendor")
	cmd.Dir = progDir
	cmd.Stdout = iotest.LogWriter(t)
	cmd.Stderr = iotest.LogWriter(t)
	require.NoError(t, cmd.Run(), "vendor dependencies")
	require.NoError(t, os.RemoveAll(filepath.Join(root, "plugin")))
	require.NoError(t, os.RemoveAll(filepath.Join(root, "dep")))
	require.NoError(t, os.RemoveAll(filepath.Join(root, "indirect-dep")))

	testPluginsAndDependencies(t, progDir)
}

// Regression test for https://github.com/khulnasoft/khulnasoft/issues/12963.
// Verifies that the language host can find plugins and dependencies
// when the Khulnasoft program is in a subdirectory of the project root.
func TestPluginsAndDependencies_subdir(t *testing.T) {
	t.Parallel()

	t.Run("moduleMode", func(t *testing.T) {
		t.Parallel()

		root := t.TempDir()
		require.NoError(t,
			fsutil.CopyFile(root, filepath.Join("testdata", "sample"), nil),
			"copy test data")

		testPluginsAndDependencies(t, filepath.Join(root, "prog-subdir", "infra"))
	})

	t.Run("vendored", func(t *testing.T) {
		t.Parallel()

		root := t.TempDir()
		require.NoError(t,
			fsutil.CopyFile(root, filepath.Join("testdata", "sample"), nil),
			"copy test data")

		progDir := filepath.Join(root, "prog-subdir", "infra")

		// Vendor the dependencies and nuke the sources
		// to ensure that the language host can only use the vendored version.
		cmd := exec.Command("go", "mod", "vendor")
		cmd.Dir = progDir
		cmd.Stdout = iotest.LogWriter(t)
		cmd.Stderr = iotest.LogWriter(t)
		require.NoError(t, cmd.Run(), "vendor dependencies")
		require.NoError(t, os.RemoveAll(filepath.Join(root, "plugin")))
		require.NoError(t, os.RemoveAll(filepath.Join(root, "dep")))
		require.NoError(t, os.RemoveAll(filepath.Join(root, "indirect-dep")))

		testPluginsAndDependencies(t, progDir)
	})

	t.Run("gowork", func(t *testing.T) {
		t.Parallel()

		root := t.TempDir()
		require.NoError(t,
			fsutil.CopyFile(root, filepath.Join("testdata", "sample"), nil),
			"copy test data")

		testPluginsAndDependencies(t, filepath.Join(root, "prog-gowork", "prog"))
	})
}

func testPluginsAndDependencies(t *testing.T, progDir string) {
	host := newLanguageHost("0.0.0.0:0", progDir, "")
	ctx := context.Background()

	t.Run("GetRequiredPackages", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		res, err := host.GetRequiredPackages(ctx, &khulnasoftrpc.GetRequiredPackagesRequest{
			Info: &khulnasoftrpc.ProgramInfo{
				RootDirectory:    progDir,
				ProgramDirectory: progDir,
				EntryPoint:       ".",
			},
		})
		require.NoError(t, err)

		require.Len(t, res.Packages, 1)
		plug := res.Packages[0]

		assert.Equal(t, "example", plug.Name, "plugin name")
		assert.Equal(t, "v1.2.3", plug.Version, "plugin version")
		assert.Equal(t, "resource", plug.Kind, "plugin kind")
		assert.Equal(t, "example.com/download", plug.Server, "plugin server")
	})

	t.Run("GetProgramDependencies", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		res, err := host.GetProgramDependencies(ctx, &khulnasoftrpc.GetProgramDependenciesRequest{
			Project:                "deprecated",
			Pwd:                    progDir,
			TransitiveDependencies: true,
			Info: &khulnasoftrpc.ProgramInfo{
				RootDirectory:    progDir,
				ProgramDirectory: progDir,
				EntryPoint:       ".",
			},
		})
		require.NoError(t, err)

		gotDeps := make(map[string]string) // name => version
		for _, dep := range res.Dependencies {
			gotDeps[dep.Name] = dep.Version
		}

		assert.Equal(t, map[string]string{
			"github.com/khulnasoft/go-dependency-testdata/plugin":          "v1.2.3",
			"github.com/khulnasoft/go-dependency-testdata/dep":             "v1.6.0",
			"github.com/khulnasoft/go-dependency-testdata/indirect-dep/v2": "v2.1.0",
		}, gotDeps)
	})
}
