// This tests that invokes cannot depend on things which are not resources.

let khulnasoft = require("../../../../../");

const dependsOn = khulnasoft.output(Promise.resolve([Promise.resolve(1)]));
khulnasoft.runtime.invokeOutput("test:index:echo", {}, { dependsOn });
