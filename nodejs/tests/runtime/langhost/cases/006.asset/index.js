// This tests simple creation of assets.

let khulnasoft = require("../../../../../");

class FileResource extends khulnasoft.CustomResource {
    constructor(name, data) {
        super("test:index:FileResource", name, {
            data: data,
        });
    }
}

new FileResource("file1", new khulnasoft.asset.FileAsset("./testdata.txt"));
