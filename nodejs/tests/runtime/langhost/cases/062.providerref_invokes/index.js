// Test the ability to invoke provider functions via RPC.

let assert = require("assert");
let khulnasoft = require("../../../../../");

(async () => {
    class Provider extends khulnasoft.ProviderResource {
        constructor(name, opts) {
            super("test", name, {}, opts);
        }
    }

    const provider = new Provider("p");
    await khulnasoft.ProviderResource.register(provider);

    let args = {
        a: "hello",
        b: true,
        c: [0.99, 42, { z: "x" }],
        id: "some-id",
        urn: "some-urn",
    };

    let result1 = await khulnasoft.runtime.invoke("test:index:echo", args, { provider });
    for (const key in args) {
        assert.deepStrictEqual(result1[key], args[key]);
    }

    let result2 = khulnasoft.runtime.invoke("test:index:echo", args, { provider, async: false });
    result2.then((v) => {
        assert.deepStrictEqual(v, args);
    });
})();
