// Test the ability to invoke provider functions via RPC.

let assert = require("assert");
let khulnasoft = require("../../../../../");
let semver = require("semver");

class Provider extends khulnasoft.ProviderResource {
    constructor(name, opts) {
        super("test", name, {}, opts);
    }
}

const provider = new Provider("p");

let args = {
    a: "hello",
    b: true,
    c: [0.99, 42, { z: "x" }],
    id: "some-id",
    urn: "some-urn",
};

let result2 = khulnasoft.runtime.invoke("test:index:echo", args, { provider });
result2.then((v) => {
    assert.deepStrictEqual(v, args);
});
