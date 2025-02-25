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

package plugin

import (
	"testing"

	"github.com/khulnasoft/sdk/go/common/resource"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDetailedDiff(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		diff     *resource.ObjectDiff
		expected map[string]PropertyDiff
	}{
		{
			name: "updates",
			diff: resource.NewPropertyMapFromMap(map[string]interface{}{
				"a": 1,
				"b": map[string]interface{}{
					"c": 2,
					"d": 3,
				},
			}).Diff(resource.NewPropertyMapFromMap(map[string]interface{}{
				"a": -1,
				"b": map[string]interface{}{
					"c": -2,
					"d": 3,
				},
			})),
			expected: map[string]PropertyDiff{
				"a": {
					Kind: DiffUpdate,
				},
				"b.c": {
					Kind: DiffUpdate,
				},
			},
		},
		{
			name: "adds and deletes",
			diff: resource.NewPropertyMapFromMap(map[string]interface{}{
				"b": map[string]interface{}{
					"c": 2,
					"d": 3,
				},
			}).Diff(resource.NewPropertyMapFromMap(map[string]interface{}{
				"a": 1,
				"b": map[string]interface{}{
					"d": 3,
				},
			})),
			expected: map[string]PropertyDiff{
				"a": {
					Kind: DiffAdd,
				},
				"b.c": {
					Kind: DiffDelete,
				},
			},
		},
		{
			name: "arrays",
			diff: resource.NewPropertyMapFromMap(map[string]interface{}{
				"a": []interface{}{
					map[string]interface{}{
						"a": 1,
						"b": []interface{}{
							2,
							3,
						},
					},
				},
			}).Diff(resource.NewPropertyMapFromMap(
				map[string]interface{}{
					"a": []interface{}{
						map[string]interface{}{
							"a": -1,
							"b": []interface{}{
								2,
							},
						},
						4,
					},
				})),
			expected: map[string]PropertyDiff{
				"a[0].a": {
					Kind: DiffUpdate,
				},
				"a[0].b[1]": {
					Kind: DiffDelete,
				},
				"a[1]": {
					Kind: DiffAdd,
				},
			},
		},
		{
			name:     "nil diff",
			diff:     nil,
			expected: map[string]PropertyDiff{},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			actual := NewDetailedDiffFromObjectDiff(c.diff, false)
			assert.Equal(t, c.expected, actual)
		})
	}
}

// Assert that UnimplementedProvider implements Provider
var _ = Provider((*UnimplementedProvider)(nil))

// Regression test for https://github.com/khulnasoft/khulnasoft/issues/14335.
// Ensure that NewDetailedDiffFromObjectDiff builds correct keys.
func TestNewDetailedDiffFromObjectDiff(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		diff          *resource.ObjectDiff
		inputDiff     bool
		expected      map[string]PropertyDiff
		expectedPaths map[string]resource.PropertyPath
	}{
		"simple add": {
			diff: &resource.ObjectDiff{
				Adds: resource.PropertyMap{
					"a": resource.NewPropertyValue(1),
				},
			},
			expected: map[string]PropertyDiff{
				"a": {Kind: DiffAdd},
			},
			expectedPaths: map[string]resource.PropertyPath{
				"a": {"a"},
			},
		},
		"simple update": {
			diff: &resource.ObjectDiff{
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"a": *resource.NewPropertyValue(1).Diff(resource.NewPropertyValue(2)),
				},
			},
			expected: map[string]PropertyDiff{
				"a": {Kind: DiffUpdate},
			},
			expectedPaths: map[string]resource.PropertyPath{
				"a": {"a"},
			},
		},
		"nested update": {
			diff: &resource.ObjectDiff{
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"a": *resource.NewObjectProperty(resource.PropertyMap{
						"b": resource.NewPropertyValue(1),
					}).Diff(resource.NewObjectProperty(resource.PropertyMap{
						"b": resource.NewPropertyValue(2),
					})),
				},
			},
			expected: map[string]PropertyDiff{
				"a.b": {Kind: DiffUpdate},
			},
			expectedPaths: map[string]resource.PropertyPath{
				"a.b": {"a", "b"},
			},
		},
		"nested update with quoted keys": {
			diff: &resource.ObjectDiff{
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"a": *resource.NewObjectProperty(resource.PropertyMap{
						"b.c":          resource.NewPropertyValue(1),
						`"quoted key"`: resource.NewPropertyValue(2),
					}).Diff(resource.NewObjectProperty(resource.PropertyMap{
						"b.c":          resource.NewPropertyValue(2),
						`"quoted key"`: resource.NewPropertyValue(3),
					})),
				},
			},
			expected: map[string]PropertyDiff{
				`a["\"quoted key\""]`: {Kind: DiffUpdate},
				`a["b.c"]`:            {Kind: DiffUpdate},
			},
			expectedPaths: map[string]resource.PropertyPath{
				`a["\"quoted key\""]`: {"a", `"quoted key"`},
				`a["b.c"]`:            {"a", "b.c"},
			},
		},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := NewDetailedDiffFromObjectDiff(tt.diff, tt.inputDiff)
			assert.Equal(t, tt.expected, result)

			for k := range result {
				path, err := resource.ParsePropertyPath(k)
				require.NoError(t, err)
				assert.Equal(t, tt.expectedPaths[k], path)
			}
		})
	}
}
