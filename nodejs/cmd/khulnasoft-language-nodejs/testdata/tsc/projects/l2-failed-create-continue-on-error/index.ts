import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as fail_on_create from "@khulnasoft/fail_on_create";
import * as simple from "@khulnasoft/simple";

const failing = new fail_on_create.Resource("failing", {value: false});
const dependent = new simple.Resource("dependent", {value: true}, {
    dependsOn: [failing],
});
const dependent_on_output = new simple.Resource("dependent_on_output", {value: failing.value});
const independent = new simple.Resource("independent", {value: true});
const double_dependency = new simple.Resource("double_dependency", {value: true}, {
    dependsOn: [
        independent,
        dependent_on_output,
    ],
});
