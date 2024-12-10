// This tests the replaceOnChanges ResourceOption.

let khulnasoft = require("../../../../../");

class MyResource extends khulnasoft.CustomResource {
    constructor(name, opts) {
        super("test:index:MyResource", name, {}, opts);
    }
}

new MyResource("testResource", { replaceOnChanges: ["foo"] });
