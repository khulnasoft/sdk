// Test the ability to invoke provider functions via RPC.

let assert = require("assert");
let khulnasoft = require("../../../../../");

class MyResource extends khulnasoft.CustomResource {
    constructor(name, args, opts) {
        super(
            "test:index:MyResource",
            name,
            Object.assign(args, {
                outprop: undefined,
            }),
            opts,
        );
    }
}

let resA = new MyResource("resA", {});
let resB = new MyResource("resB", {}, { dependsOn: [resA] });
let resC = new MyResource("resC", {
    propA: resA.outprop,
    propB: resB.outprop,
    propC: "foo",
});
let resD = new MyResource("resD", {
    propA: khulnasoft.all([resA.outprop, resB.outprop]).apply(([a, b]) => `${a} ${b}`),
    propB: resC.outprop,
    propC: "bar",
});
let resE = new MyResource(
    "resE",
    {
        propA: resC.outprop,
        propB: khulnasoft.all([resA.outprop, resB.outprop]).apply(([a, b]) => `${a} ${b}`),
        propC: "baz",
    },
    { dependsOn: [resD] },
);
