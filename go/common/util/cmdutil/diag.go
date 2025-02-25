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

package cmdutil

import (
	"fmt"
	"os"
	"sync"

	"github.com/khulnasoft/sdk/go/common/diag"
	"github.com/khulnasoft/sdk/go/common/diag/colors"
	"github.com/khulnasoft/sdk/go/common/util/contract"
)

var (
	snkMutex sync.Mutex
	snk      diag.Sink
)

// By default we'll attempt to figure out if we should have colors or not. This can be overridden
// for any command by passing --color=... at the command line.
var globalColorization *colors.Colorization

// GetGlobalColorization gets the global setting for how things should be colored.
// This is helpful for the parts of our stack that do not take a DisplayOptions struct.
func GetGlobalColorization() colors.Colorization {
	if globalColorization != nil {
		// User has set an explicit colorization preference.  We'll respect whatever they asked for,
		// no matter what.
		return *globalColorization
	}

	// Colorization is set to 'auto' (either explicit set to that by the user, or not set at all).
	// Figure out the best thing to do here.

	// If the external environment has requested no colors, then turn off all colors when in 'auto' mode.
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		return colors.Never
	}

	// Disable colors if we're not in an interactive session (i.e. we're redirecting stdout).  This
	// will just inject color tags into the stream which are not desirable here.
	if !InteractiveTerminal() {
		return colors.Never
	}

	// Things otherwise look good.  Turn on colors.
	return colors.Always
}

// SetGlobalColorization sets the global setting for how things should be colored.
// This is helpful for the parts of our stack that do not take a DisplayOptions struct.
func SetGlobalColorization(value string) error {
	switch value {
	case "auto":
		globalColorization = nil
	case "always":
		c := colors.Always
		globalColorization = &c
	case "never":
		c := colors.Never
		globalColorization = &c
	case "raw":
		c := colors.Raw
		globalColorization = &c
	default:
		return fmt.Errorf("unsupported color option: '%s'.  Supported values are: auto, always, never, raw", value)
	}

	return nil
}

// Diag lazily allocates a sink to be used if we can't create a compiler.
func Diag() diag.Sink {
	snkMutex.Lock()
	defer snkMutex.Unlock()
	if snk == nil {
		snk = diag.DefaultSink(os.Stdout, os.Stderr, diag.FormatOptions{
			Color: GetGlobalColorization(),
		})
	}
	return snk
}

// InitDiag forces initialization of the diagnostics sink with the given options.
func InitDiag(opts diag.FormatOptions) {
	contract.Assertf(snk == nil, "Cannot initialize diagnostics sink more than once")
	snk = diag.DefaultSink(os.Stdout, os.Stderr, opts)
}
