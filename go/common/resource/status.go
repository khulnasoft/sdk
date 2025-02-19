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

package resource

// Status is returned when an error has occurred during a resource provider operation.  It indicates whether the
// operation could be rolled back cleanly (OK).  If not, it means the resource was left in an indeterminate state.
type Status int

const (
	StatusOK Status = iota
	StatusPartialFailure
	StatusUnknown
)
