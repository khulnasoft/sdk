import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as simple_invoke from "@khulnasoft/simple-invoke";

const explicitProvider = new simple_invoke.Provider("explicitProvider", {});
const first = new simple_invoke.StringResource("first", {text: "first hello"});
const data = simple_invoke.myInvokeOutput({
    value: "hello",
}, {
    dependsOn: [first],
});
const second = new simple_invoke.StringResource("second", {text: data.result});
export const hello = data.result;
