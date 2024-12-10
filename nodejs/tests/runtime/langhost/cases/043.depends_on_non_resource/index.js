// This tests that resources cannot depend on things which are not resources.

let khulnasoft = require("../../../../../");

class MyResource extends khulnasoft.CustomResource {
    constructor(name, deps) {
        super("test:index:MyResource", name, {}, deps);
    }
}

new MyResource("testResource", { dependsOn: khulnasoft.output(Promise.resolve([Promise.resolve(1)])) });
