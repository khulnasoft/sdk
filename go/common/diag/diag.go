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

package diag

import (
	"github.com/khulnasoft/sdk/go/common/resource"
)

// ID is a unique diagnostics identifier.
type ID int

// Diag is an instance of an error or warning generated by the compiler.
type Diag struct {
	URN     resource.URN // Resource this diagnostics is associated with.  Empty if not associated with any resource.
	ID      ID           // a unique identifier for this diagnostic.
	Message string       // a human-friendly message for this diagnostic.
	Raw     bool         // true if this diagnostic should not be formatted when displayed.

	// An ID used to collate a stream of conceptually sequential messages.  0 means that the message
	// is not part of any sequential message stream.
	StreamID int32
}

// Message returns an anonymous diagnostic message without any source or ID information.
func Message(urn resource.URN, msg string) *Diag {
	return &Diag{URN: urn, Message: msg}
}

// RawMessage returns an anonymous diagnostic message without any source or ID information that will not be rendered
// with Sprintf.
func RawMessage(urn resource.URN, msg string) *Diag {
	return &Diag{URN: urn, Message: msg, Raw: true}
}

// StreamMessage returns an anonymous diagnostic message without any source or ID information that
// is associated with the given stream ID.  Displays can use this ID to combine all the messages
// from a single stream into an entire message, while still rendering the pieces as they come in.
func StreamMessage(urn resource.URN, msg string, streamID int32) *Diag {
	return &Diag{URN: urn, Message: msg, Raw: true, StreamID: streamID}
}
