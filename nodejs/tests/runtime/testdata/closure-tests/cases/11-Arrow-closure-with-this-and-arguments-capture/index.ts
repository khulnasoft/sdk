// Copyright 2024-2024, KhulnaSoft, Ltd.
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

export const description = "Arrow closure with this and arguments capture";

// @ts-ignore
export const func = (function () { return () => { console.log(this + arguments); } }).apply(this, [0, 1])

export const error = `Error serializing function '<anonymous>'

function '<anonymous>': which could not be serialized because
  arrow function captured 'this'. Assign 'this' to another name outside function and capture that.

Function code:
  () => { console.log(this + arguments); }
`