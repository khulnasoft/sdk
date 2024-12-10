let assert = require("assert");
let khulnasoft = require("../../../../../");

class MyResource extends khulnasoft.CustomResource {
    constructor(name, args, opts) {
        super("test:index:MyResource", name, args, opts);
    }
}

//            cust1
//                \
//                 cust2

let cust1 = new MyResource("cust1", {});
let cust2 = new MyResource("cust2", { parentId: cust1.id }, { parent: cust1 });
