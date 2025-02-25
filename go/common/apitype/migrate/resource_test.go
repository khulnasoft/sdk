// Copyright 2016-2018, KhulnaSoft, Ltd.
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

package migrate

import (
	"testing"

	"github.com/khulnasoft/sdk/go/common/apitype"
	"github.com/khulnasoft/sdk/go/common/resource"
	"github.com/khulnasoft/sdk/go/common/tokens"
	"github.com/stretchr/testify/assert"
)

func TestV1ToV2(t *testing.T) {
	t.Parallel()

	v1 := apitype.ResourceV1{
		URN:    resource.URN("foo"),
		Custom: true,
		Delete: true,
		ID:     resource.ID("bar"),
		Type:   tokens.Type("special"),
		Inputs: map[string]interface{}{
			"foo_in": "baz",
		},
		Defaults: map[string]interface{}{
			"foo_default": "stuff",
		},
		Outputs: map[string]interface{}{
			"foo_out": "out",
		},
		Parent:  resource.URN("parent"),
		Protect: true,
		Dependencies: []resource.URN{
			resource.URN("dep1"),
			resource.URN("dep2"),
		},
	}

	v2 := UpToResourceV2(v1)
	assert.Equal(t, resource.URN("foo"), v2.URN)
	assert.True(t, v2.Custom)
	assert.True(t, v2.Delete)
	assert.Equal(t, resource.ID("bar"), v2.ID)
	assert.Equal(t, tokens.Type("special"), v2.Type)
	assert.Equal(t, map[string]interface{}{
		"foo_in": "baz",
	}, v2.Inputs)
	assert.Equal(t, map[string]interface{}{
		"foo_out": "out",
	}, v2.Outputs)
	assert.Equal(t, resource.URN("parent"), v2.Parent)
	assert.True(t, v2.Protect)
	assert.False(t, v2.External)
	assert.Equal(t, []resource.URN{
		resource.URN("dep1"),
		resource.URN("dep2"),
	}, v2.Dependencies)
	assert.Empty(t, v2.Provider)
}

func TestV2ToV3(t *testing.T) {
	t.Parallel()

	v2 := apitype.ResourceV2{
		URN:    resource.URN("foo"),
		Custom: true,
		Delete: true,
		ID:     resource.ID("bar"),
		Type:   tokens.Type("special"),
		Inputs: map[string]interface{}{
			"foo_in": "baz",
		},
		Outputs: map[string]interface{}{
			"foo_out": "out",
		},
		Parent:   resource.URN("parent"),
		Protect:  true,
		External: false,
		Dependencies: []resource.URN{
			resource.URN("dep1"),
			resource.URN("dep2"),
		},
		Provider:   "provider",
		InitErrors: []string{"error"},
	}

	v3 := UpToResourceV3(v2)

	assert.Equal(t, resource.URN("foo"), v3.URN)
	assert.True(t, v3.Custom)
	assert.True(t, v3.Delete)
	assert.Equal(t, resource.ID("bar"), v3.ID)
	assert.Equal(t, tokens.Type("special"), v3.Type)
	assert.Equal(t, map[string]interface{}{
		"foo_in": "baz",
	}, v3.Inputs)
	assert.Equal(t, map[string]interface{}{
		"foo_out": "out",
	}, v3.Outputs)
	assert.Equal(t, resource.URN("parent"), v3.Parent)
	assert.True(t, v3.Protect)
	assert.False(t, v3.External)
	assert.Equal(t, []resource.URN{
		resource.URN("dep1"),
		resource.URN("dep2"),
	}, v3.Dependencies)
	assert.Equal(t, "provider", v3.Provider)
	assert.Equal(t, []string{"error"}, v3.InitErrors)
	assert.Nil(t, v3.Created)
	assert.Nil(t, v3.Modified)
}
