// This tests the basic creation of a single propertyless resource.

let khulnasoft = require("../../../../../");

class MyResource extends khulnasoft.CustomResource {
    constructor(name, opts) {
        super("test:index:MyResource", name, {}, opts);
    }
}

new MyResource("testResource1", { import: "testID" });
