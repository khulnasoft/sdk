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

export const description = "Three level inheritance with symbols";

const sym = Symbol();

class A {
    public n: number;
    constructor(n: number) {
        this.n = n;
        console.log("AConstruction");
    }
    public [sym](x: number) { return x; }
}
class B extends A {
    constructor(n: number) {
        super(n + 1);
        console.log("BConstructor");
    }
    // @ts-ignore
    public [sym](n: number) { return 1 + super[sym](n + 1); }
}
class C extends B {
    constructor(n: number) {
        super(n * 2);
        console.log("CConstructor");
    }
    // @ts-ignore
    public [sym](n: number) { return 2 * super[sym](n * 2); }
}

export const func = () => C;
