// This tests that providers is passed to RegisterResource for remote components, whether specified via
// `provider` or `providers`.

import khulnasoft from "../../../../../index.js";

class Provider extends khulnasoft.ProviderResource {
    constructor(name, opts) {
        super("test", name, {}, opts);
    }
}

class FooProvider extends khulnasoft.ProviderResource {
    constructor(name, opts) {
        super("foo", name, {}, opts);
    }
}

class RemoteComponent extends khulnasoft.ComponentResource {
    constructor(name, opts) {
        super("test:index:Component", name, {}, opts, true /*remote*/);
    }
}

const myprovider = new Provider("myprovider");

new RemoteComponent("singular", { provider: myprovider });
new RemoteComponent("map", { providers: { test: myprovider } });
new RemoteComponent("array", { providers: [myprovider] });

const fooprovider = new FooProvider("fooprovider");

new RemoteComponent("foo-singular", { provider: fooprovider });
new RemoteComponent("foo-map", { providers: { foo: fooprovider } });
new RemoteComponent("foo-array", { providers: [fooprovider] });
