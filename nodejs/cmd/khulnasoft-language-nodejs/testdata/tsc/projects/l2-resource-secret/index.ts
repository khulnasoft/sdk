import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as secret from "@khulnasoft/secret";

const res = new secret.Resource("res", {
    "private": "closed",
    "public": "open",
    privateData: {
        "private": "closed",
        "public": "open",
    },
    publicData: {
        "private": "closed",
        "public": "open",
    },
});
