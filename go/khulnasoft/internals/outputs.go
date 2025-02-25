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

package internals

import (
	"context"
	_ "unsafe" // unsafe is needed to use go:linkname

	"github.com/khulnasoft/sdk/go/khulnasoft"
)

//go:linkname awaitWithContext github.com/khulnasoft/sdk/go/khulnasoft.awaitWithContext
func awaitWithContext(ctx context.Context, o khulnasoft.Output) (interface{}, bool, bool, []khulnasoft.Resource, error)

// UnsafeAwaitOutputResult is an output from a Khulnasoft function or resource that has been resolved.
//
// This is a low level API and should be used with care.
type UnsafeAwaitOutputResult struct {
	Value        interface{}           // The value of the output. If unknown (in a dry-run), the value will be nil.
	Known        bool                  // True if the value is known.
	Secret       bool                  // True if the value is a secret.
	Dependencies []khulnasoft.Resource // The resources that this output depends on.
}

// UnsafeAwaitOutput blocks until the output is resolved and returns the resolved value and
// metadata.
//
// This is a low level API and should be used with care.
func UnsafeAwaitOutput(ctx context.Context, o khulnasoft.Output) (UnsafeAwaitOutputResult, error) {
	value, known, secret, deps, err := awaitWithContext(ctx, o)

	return UnsafeAwaitOutputResult{
		Value:        value,
		Known:        known,
		Secret:       secret,
		Dependencies: deps,
	}, err
}
