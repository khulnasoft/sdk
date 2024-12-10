let assert = require("assert");
let khulnasoft = require("../../../../../");

class MyCustomResource extends khulnasoft.CustomResource {
    constructor(name, opts) {
        super("test:index:MyCustomResource", name, {}, opts);
    }
}

class MyComponentResource extends khulnasoft.ComponentResource {
    constructor(name, opts) {
        super("test:index:MyComponentResource", name, {}, opts);
    }
}

const custom = new MyCustomResource("custom");
const component = new MyComponentResource("component");
