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

let first = new MyComponentResource("first");
// The test checks that "second" depends on "firstChild".
let firstChild = new MyCustomResource("firstChild", {
    parent: first,
});
// The test looks for this resource named "second".
let second = new MyComponentResource("second", {
    parent: first,
    dependsOn: first,
});
let myresource = new MyCustomResource("myresource", {
    parent: second,
});
