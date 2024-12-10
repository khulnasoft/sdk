// Ensure that certain runtime settings are available.

const assert = require("assert");
const khulnasoft = require("../../../../../");

assert.strictEqual(khulnasoft.getProject(), "runtimeSettingsProject");
assert.strictEqual(khulnasoft.getStack(), "runtimeSettingsStack");

const config = new khulnasoft.Config("myBag");
assert.strictEqual(config.getNumber("A"), 42);
assert.strictEqual(config.requireNumber("A"), 42);
assert.strictEqual(config.get("bbbb"), "a string o' b's");
assert.strictEqual(config.require("bbbb"), "a string o' b's");
assert.strictEqual(config.get("missingC"), undefined);

// ensure the older format works as well
const configOld = new khulnasoft.Config("myBag:config");
assert.strictEqual(configOld.getNumber("A"), 42);
assert.strictEqual(configOld.requireNumber("A"), 42);
assert.strictEqual(configOld.get("bbbb"), "a string o' b's");
assert.strictEqual(configOld.require("bbbb"), "a string o' b's");
assert.strictEqual(configOld.get("missingC"), undefined);
