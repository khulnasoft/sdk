// This tests the runtime's ability to be loaded side-by-side with another copy of the same runtime library.
// This is a hard and subtle problem because the runtime is configured with a bunch of state, like whether
// we are doing a dry-run and, more importantly, RPC addresses to communicate with the engine.  Normally we
// go through the startup shim to configure all of these things, but when the second copy gets loaded we don't.
// Subsequent copies of the runtime are able to configure themselves by using environment variables.

let assert = require("assert");
let path = require("path");

const sdkPath = "../../../../../";

// Load the first copy:
let khulnasoft1 = require(sdkPath);

// Now delete the entries in the require cache, and load up the second copy:
const resolvedSdkPath = path.dirname(require.resolve(sdkPath));
Object.keys(require.cache).forEach((path) => {
    if (path.startsWith(resolvedSdkPath)) {
        delete require.cache[path];
    }
});
let khulnasoft2 = require(sdkPath);

// Make sure they are different:
assert(khulnasoft1 !== khulnasoft2, "khulnasoft1 !== khulnasoft2");
assert(khulnasoft1.runtime !== khulnasoft2.runtime, "khulnasoft1.runtime !== khulnasoft2.runtime");

// Check that various settings are equal:
assert.strictEqual(
    khulnasoft1.runtime.isDryRun(),
    khulnasoft2.runtime.isDryRun(),
    "khulnasoft1.runtime.isDryRun() !== khulnasoft2.runtime.isDryRun()",
);
assert.strictEqual(
    khulnasoft1.runtime.getProject(),
    khulnasoft2.runtime.getProject(),
    "khulnasoft1.runtime.getProject() !== khulnasoft2.runtime.getProject()",
);
assert.strictEqual(
    khulnasoft1.runtime.getStack(),
    khulnasoft2.runtime.getStack(),
    "khulnasoft1.runtime.getStack() !== khulnasoft2.runtime.getStack()",
);
assert.deepStrictEqual(
    khulnasoft1.runtime.allConfig(),
    khulnasoft2.runtime.allConfig(),
    "khulnasoft1.runtime.allConfig() !== khulnasoft2.runtime.getStack()",
);

// Check that the two runtimes agree on the stack resource
let stack1 = khulnasoft1.runtime.getStackResource();
let stack2 = khulnasoft2.runtime.getStackResource();
assert.strictEqual(stack1, stack2, "khulnasoft1.runtime.getStackResource() !== khulnasoft2.runtime.getStackResource()");

// allConfig should have caught this, but let's check individual config values too.
let cfg1 = new khulnasoft1.Config("sxs");
let cfg2 = new khulnasoft2.Config("sxs");
assert.strictEqual(cfg1.get("message"), cfg2.get("message"));

// Try and set a stack transformation
function transform1(args) {
    args.props["runtime1"] = 1;
    return { props: args.props, opts: args.opts };
}
function transform2(args) {
    args.props["runtime2"] = 2;
    return { props: args.props, opts: args.opts };
}

khulnasoft1.runtime.registerStackTransformation(transform1);
khulnasoft2.runtime.registerStackTransformation(transform2);

// Now do some useful things that require RPC connections:
khulnasoft1.log.info("logging via Khulnasoft1 works!");
khulnasoft2.log.info("logging via Khulnasoft2 works too!");
let res1 = new khulnasoft1.CustomResource("test:x:resource", "p1p1p1");
res1.urn.apply((urn) => assert.strictEqual(urn, "test:x:resource::p1p1p1"));
let res2 = new khulnasoft2.CustomResource("test:y:resource", "p2p2p2");
res2.urn.apply((urn) => assert.strictEqual(urn, "test:y:resource::p2p2p2"));

// Both resources should have the stack transforms applied
res1.runtime1.apply((value) => assert.strictEqual(value, 1));
res1.runtime2.apply((value) => assert.strictEqual(value, 2));
res2.runtime1.apply((value) => assert.strictEqual(value, 1));
res2.runtime2.apply((value) => assert.strictEqual(value, 2));

khulnasoft1.runtime.registerResourcePackage("test1", {
    version: "0.0.1",
});
khulnasoft2.runtime.registerResourcePackage("test2", {
    version: "0.0.2",
});
let test1khulnasoft1 = khulnasoft1.runtime.getResourcePackage("test1");
assert.strictEqual(test1khulnasoft1.version, "0.0.1");
let test1khulnasoft2 = khulnasoft2.runtime.getResourcePackage("test1");
assert.strictEqual(test1khulnasoft2.version, "0.0.1");
let test2khulnasoft1 = khulnasoft1.runtime.getResourcePackage("test2");
assert.strictEqual(test2khulnasoft1.version, "0.0.2");
let test2khulnasoft2 = khulnasoft2.runtime.getResourcePackage("test2");
assert.strictEqual(test2khulnasoft2.version, "0.0.2");

// Check that we can set mocks successfully
// We don't need an actual test monitor here, just something to set and get.
class Mocks {
    newResource(args) {
        return {
            id: args.inputs.name + "_id",
            state: {
                ...args.inputs,
            },
        };
    }
    call(args) {
        return args;
    }
}

let mocks = new Mocks();

khulnasoft1.runtime.setMocks(mocks);
assert.strictEqual(khulnasoft1.runtime.getMonitor().mocks, mocks);
assert.strictEqual(khulnasoft2.runtime.getMonitor().mocks, mocks);
assert.strictEqual(khulnasoft1.runtime.getMonitor(), khulnasoft2.runtime.getMonitor());
assert(khulnasoft1.runtime.hasMonitor());
assert(khulnasoft2.runtime.hasMonitor());

khulnasoft1.runtime._reset();
khulnasoft2.runtime._reset();
