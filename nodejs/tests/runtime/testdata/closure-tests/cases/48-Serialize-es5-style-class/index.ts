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

export const description = "Serialize es5-style class";

const outer: any = { o: 1 };
const array = [outer];
outer.b = array;
const C = (function () {
    function C() { }
    C.prototype.m = function () { return this.n() + outer; };
    C.prototype.n = function () { return array; };
    (<any>C).m = function () { return this.n(); };
    return C;
}());

export const func = () => C;
