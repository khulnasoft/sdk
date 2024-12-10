// This tests the basic creation of a single propertyless resource.

let khulnasoft = require("../../../../../");

class MyResource extends khulnasoft.CustomResource {
    constructor(name) {
        super("test:index:MyResource", name);
    }
}

new MyResource("testResource1");
