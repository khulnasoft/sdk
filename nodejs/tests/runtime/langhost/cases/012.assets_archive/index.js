// This tests the basic creation of a single resource with an assets archive property.

let khulnasoft = require("../../../../../");

class MyResource extends khulnasoft.CustomResource {
    constructor(name) {
        let archive = new khulnasoft.asset.AssetArchive({
            asset: new khulnasoft.asset.StringAsset("foo"),
            archive: new khulnasoft.asset.AssetArchive({}),
        });
        let archiveP = Promise.resolve(
            new khulnasoft.asset.AssetArchive({
                foo: new khulnasoft.asset.StringAsset("bar"),
            }),
        );
        let assetP = Promise.resolve(new khulnasoft.asset.StringAsset("baz"));
        super("test:index:MyResource", name, {
            archive: archive,
            archiveP: archiveP,
            assetP: assetP,
        });
    }
}

new MyResource("testResource1");
