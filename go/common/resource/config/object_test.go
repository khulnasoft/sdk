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

package config

import (
	"context"
	"testing"

	"github.com/khulnasoft/sdk/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmptyObject(t *testing.T) {
	t.Parallel()

	// Test that an empty object can be converted to a property value
	// without error.
	o := object{}
	crypter := nopCrypter{}
	v, err := o.toDecryptedPropertyValue(context.Background(), crypter)
	require.NoError(t, err)
	assert.Equal(t, resource.NewNullProperty(), v)
}
