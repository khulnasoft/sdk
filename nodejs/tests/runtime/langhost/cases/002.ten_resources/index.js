// This tests the creation of ten propertyless resources.

let khulnasoft = require("../../../../../");

class MyResource extends khulnasoft.CustomResource {
    constructor(name) {
        super("test:index:MyResource", name);
    }
}

for (let i = 0; i < 10; i++) {
    new MyResource("testResource" + i);
}
