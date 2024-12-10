let khulnasoft = require("../../../../../");

khulnasoft.log.info("info message");
khulnasoft.log.warn("warning message");
khulnasoft.log.error("error message");

class FakeResource extends khulnasoft.CustomResource {
    constructor(name) {
        super("test:index:FakeResource", name);
    }
}

const res = new FakeResource("test");
khulnasoft.log.info("attached to resource", res);
khulnasoft.log.info("with streamid", res, 42);
