let khulnasoft = require("../../../../../");

class MyResource extends khulnasoft.CustomResource {
    constructor(name, version) {
        super("test:index:MyResource", name, {}, { version: version });
    }
}

new MyResource("testResource", "0.19.1");
new MyResource("testResource2", "0.19.2");
new MyResource("testResource3");

khulnasoft.runtime.invoke("invoke:index:doit", {}, { version: "0.19.1", async: false });
khulnasoft.runtime.invoke("invoke:index:doit_v2", {}, { version: "0.19.2", async: false });
khulnasoft.runtime.invoke("invoke:index:doit_noversion", { async: false });

khulnasoft.runtime.invoke("invoke:index:doit", {}, { version: "0.19.1" });
khulnasoft.runtime.invoke("invoke:index:doit_v2", {}, { version: "0.19.2" });
khulnasoft.runtime.invoke("invoke:index:doit_noversion", {});

new khulnasoft.CustomResource("test:index:ReadResource", "foo", {}, { id: "readme", version: "0.20.0" });
new khulnasoft.CustomResource("test:index:ReadResource", "foo_noversion", {}, { id: "readme" });
