// This tests the creation of ten propertyless resources.

let khulnasoft = require("../../../../../");

class MyResource extends khulnasoft.CustomResource {
    constructor(name, deps) {
        super("test:index:MyResource", name, {}, deps);
    }
}

let all = [];
let last = undefined;
for (let i = 0; i < 10; i++) {
    // Test all permutations of accepted dependsOn input:
    //     - undefined
    //     - Resource
    //     - Resource[]
    //     - Promise<Resource>
    //     - Output<Resource>
    //     - Promise<Resource[]>
    //     - Output<Resource[]>
    //     - Promise<Promise<Resource>[]>
    //     - Promise<Output<Resource>[]>
    //     - Output<Promise<Resource>[]>
    //     - Output<Output<Resource>[]>
    let r0 = new MyResource("testResource" + i * 10, { dependsOn: last });
    let r1 = new MyResource("testResource" + i * 10 + 1, { dependsOn: all });
    let r2 = new MyResource("testResource" + i * 10 + 2, { dependsOn: Promise.resolve(last) });
    let r3 = new MyResource("testResource" + i * 10 + 3, { dependsOn: khulnasoft.Output.create(last) });
    let r4 = new MyResource("testResource" + i * 10 + 4, { dependsOn: Promise.resolve(all) });
    let r5 = new MyResource("testResource" + i * 10 + 5, { dependsOn: khulnasoft.Output.create(all) });
    let r6 = new MyResource("testResource" + i * 10 + 6, {
        dependsOn: Promise.resolve(all.map((a) => Promise.resolve(a))),
    });
    let r7 = new MyResource("testResource" + i * 10 + 7, {
        dependsOn: Promise.resolve(all.map((a) => khulnasoft.Output.create(a))),
    });
    let r8 = new MyResource("testResource" + i * 10 + 8, {
        dependsOn: khulnasoft.Output.create(all).apply((a) => Promise.resolve(a)),
    });
    let r9 = new MyResource("testResource" + i * 10 + 9, {
        dependsOn: khulnasoft.Output.create(all).apply((a) => khulnasoft.Output.create(a)),
    });
    all = all.concat([r0, r1, r2, r3, r4, r5, r6, r7, r8, r9]);
    last = r0;
}
