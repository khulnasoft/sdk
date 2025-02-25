// *** WARNING: this file was generated by khulnasoft-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class Provider extends khulnasoft.ProviderResource {
    /** @internal */
    public static readonly __khulnasoftType = 'config-grpc';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Khulnasoft SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__khulnasoftType'] === "khulnasoft:providers:" + Provider.__khulnasoftType;
    }

    public readonly secretString1!: khulnasoft.Output<string | undefined>;
    public readonly secretString2!: khulnasoft.Output<string | undefined>;
    public readonly secretString3!: khulnasoft.Output<string | undefined>;
    public readonly string1!: khulnasoft.Output<string | undefined>;
    public readonly string2!: khulnasoft.Output<string | undefined>;
    public readonly string3!: khulnasoft.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: khulnasoft.ResourceOptions) {
        let resourceInputs: khulnasoft.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["bool1"] = khulnasoft.output(args ? args.bool1 : undefined).apply(JSON.stringify);
            resourceInputs["bool2"] = khulnasoft.output(args ? args.bool2 : undefined).apply(JSON.stringify);
            resourceInputs["bool3"] = khulnasoft.output(args ? args.bool3 : undefined).apply(JSON.stringify);
            resourceInputs["int1"] = khulnasoft.output(args ? args.int1 : undefined).apply(JSON.stringify);
            resourceInputs["int2"] = khulnasoft.output(args ? args.int2 : undefined).apply(JSON.stringify);
            resourceInputs["int3"] = khulnasoft.output(args ? args.int3 : undefined).apply(JSON.stringify);
            resourceInputs["listBool1"] = khulnasoft.output(args ? args.listBool1 : undefined).apply(JSON.stringify);
            resourceInputs["listBool2"] = khulnasoft.output(args ? args.listBool2 : undefined).apply(JSON.stringify);
            resourceInputs["listBool3"] = khulnasoft.output(args ? args.listBool3 : undefined).apply(JSON.stringify);
            resourceInputs["listInt1"] = khulnasoft.output(args ? args.listInt1 : undefined).apply(JSON.stringify);
            resourceInputs["listInt2"] = khulnasoft.output(args ? args.listInt2 : undefined).apply(JSON.stringify);
            resourceInputs["listInt3"] = khulnasoft.output(args ? args.listInt3 : undefined).apply(JSON.stringify);
            resourceInputs["listNum1"] = khulnasoft.output(args ? args.listNum1 : undefined).apply(JSON.stringify);
            resourceInputs["listNum2"] = khulnasoft.output(args ? args.listNum2 : undefined).apply(JSON.stringify);
            resourceInputs["listNum3"] = khulnasoft.output(args ? args.listNum3 : undefined).apply(JSON.stringify);
            resourceInputs["listSecretBool1"] = khulnasoft.output(args?.listSecretBool1 ? khulnasoft.secret(args.listSecretBool1) : undefined).apply(JSON.stringify);
            resourceInputs["listSecretBool2"] = khulnasoft.output(args?.listSecretBool2 ? khulnasoft.secret(args.listSecretBool2) : undefined).apply(JSON.stringify);
            resourceInputs["listSecretBool3"] = khulnasoft.output(args?.listSecretBool3 ? khulnasoft.secret(args.listSecretBool3) : undefined).apply(JSON.stringify);
            resourceInputs["listSecretInt1"] = khulnasoft.output(args?.listSecretInt1 ? khulnasoft.secret(args.listSecretInt1) : undefined).apply(JSON.stringify);
            resourceInputs["listSecretInt2"] = khulnasoft.output(args?.listSecretInt2 ? khulnasoft.secret(args.listSecretInt2) : undefined).apply(JSON.stringify);
            resourceInputs["listSecretInt3"] = khulnasoft.output(args?.listSecretInt3 ? khulnasoft.secret(args.listSecretInt3) : undefined).apply(JSON.stringify);
            resourceInputs["listSecretNum1"] = khulnasoft.output(args?.listSecretNum1 ? khulnasoft.secret(args.listSecretNum1) : undefined).apply(JSON.stringify);
            resourceInputs["listSecretNum2"] = khulnasoft.output(args?.listSecretNum2 ? khulnasoft.secret(args.listSecretNum2) : undefined).apply(JSON.stringify);
            resourceInputs["listSecretNum3"] = khulnasoft.output(args?.listSecretNum3 ? khulnasoft.secret(args.listSecretNum3) : undefined).apply(JSON.stringify);
            resourceInputs["listSecretString1"] = khulnasoft.output(args?.listSecretString1 ? khulnasoft.secret(args.listSecretString1) : undefined).apply(JSON.stringify);
            resourceInputs["listSecretString2"] = khulnasoft.output(args?.listSecretString2 ? khulnasoft.secret(args.listSecretString2) : undefined).apply(JSON.stringify);
            resourceInputs["listSecretString3"] = khulnasoft.output(args?.listSecretString3 ? khulnasoft.secret(args.listSecretString3) : undefined).apply(JSON.stringify);
            resourceInputs["listString1"] = khulnasoft.output(args ? args.listString1 : undefined).apply(JSON.stringify);
            resourceInputs["listString2"] = khulnasoft.output(args ? args.listString2 : undefined).apply(JSON.stringify);
            resourceInputs["listString3"] = khulnasoft.output(args ? args.listString3 : undefined).apply(JSON.stringify);
            resourceInputs["mapBool1"] = khulnasoft.output(args ? args.mapBool1 : undefined).apply(JSON.stringify);
            resourceInputs["mapBool2"] = khulnasoft.output(args ? args.mapBool2 : undefined).apply(JSON.stringify);
            resourceInputs["mapBool3"] = khulnasoft.output(args ? args.mapBool3 : undefined).apply(JSON.stringify);
            resourceInputs["mapInt1"] = khulnasoft.output(args ? args.mapInt1 : undefined).apply(JSON.stringify);
            resourceInputs["mapInt2"] = khulnasoft.output(args ? args.mapInt2 : undefined).apply(JSON.stringify);
            resourceInputs["mapInt3"] = khulnasoft.output(args ? args.mapInt3 : undefined).apply(JSON.stringify);
            resourceInputs["mapNum1"] = khulnasoft.output(args ? args.mapNum1 : undefined).apply(JSON.stringify);
            resourceInputs["mapNum2"] = khulnasoft.output(args ? args.mapNum2 : undefined).apply(JSON.stringify);
            resourceInputs["mapNum3"] = khulnasoft.output(args ? args.mapNum3 : undefined).apply(JSON.stringify);
            resourceInputs["mapSecretBool1"] = khulnasoft.output(args?.mapSecretBool1 ? khulnasoft.secret(args.mapSecretBool1) : undefined).apply(JSON.stringify);
            resourceInputs["mapSecretBool2"] = khulnasoft.output(args?.mapSecretBool2 ? khulnasoft.secret(args.mapSecretBool2) : undefined).apply(JSON.stringify);
            resourceInputs["mapSecretBool3"] = khulnasoft.output(args?.mapSecretBool3 ? khulnasoft.secret(args.mapSecretBool3) : undefined).apply(JSON.stringify);
            resourceInputs["mapSecretInt1"] = khulnasoft.output(args?.mapSecretInt1 ? khulnasoft.secret(args.mapSecretInt1) : undefined).apply(JSON.stringify);
            resourceInputs["mapSecretInt2"] = khulnasoft.output(args?.mapSecretInt2 ? khulnasoft.secret(args.mapSecretInt2) : undefined).apply(JSON.stringify);
            resourceInputs["mapSecretInt3"] = khulnasoft.output(args?.mapSecretInt3 ? khulnasoft.secret(args.mapSecretInt3) : undefined).apply(JSON.stringify);
            resourceInputs["mapSecretNum1"] = khulnasoft.output(args?.mapSecretNum1 ? khulnasoft.secret(args.mapSecretNum1) : undefined).apply(JSON.stringify);
            resourceInputs["mapSecretNum2"] = khulnasoft.output(args?.mapSecretNum2 ? khulnasoft.secret(args.mapSecretNum2) : undefined).apply(JSON.stringify);
            resourceInputs["mapSecretNum3"] = khulnasoft.output(args?.mapSecretNum3 ? khulnasoft.secret(args.mapSecretNum3) : undefined).apply(JSON.stringify);
            resourceInputs["mapSecretString1"] = khulnasoft.output(args?.mapSecretString1 ? khulnasoft.secret(args.mapSecretString1) : undefined).apply(JSON.stringify);
            resourceInputs["mapSecretString2"] = khulnasoft.output(args?.mapSecretString2 ? khulnasoft.secret(args.mapSecretString2) : undefined).apply(JSON.stringify);
            resourceInputs["mapSecretString3"] = khulnasoft.output(args?.mapSecretString3 ? khulnasoft.secret(args.mapSecretString3) : undefined).apply(JSON.stringify);
            resourceInputs["mapString1"] = khulnasoft.output(args ? args.mapString1 : undefined).apply(JSON.stringify);
            resourceInputs["mapString2"] = khulnasoft.output(args ? args.mapString2 : undefined).apply(JSON.stringify);
            resourceInputs["mapString3"] = khulnasoft.output(args ? args.mapString3 : undefined).apply(JSON.stringify);
            resourceInputs["num1"] = khulnasoft.output(args ? args.num1 : undefined).apply(JSON.stringify);
            resourceInputs["num2"] = khulnasoft.output(args ? args.num2 : undefined).apply(JSON.stringify);
            resourceInputs["num3"] = khulnasoft.output(args ? args.num3 : undefined).apply(JSON.stringify);
            resourceInputs["objBool1"] = khulnasoft.output(args ? args.objBool1 : undefined).apply(JSON.stringify);
            resourceInputs["objBool2"] = khulnasoft.output(args ? args.objBool2 : undefined).apply(JSON.stringify);
            resourceInputs["objBool3"] = khulnasoft.output(args ? args.objBool3 : undefined).apply(JSON.stringify);
            resourceInputs["objInt1"] = khulnasoft.output(args ? args.objInt1 : undefined).apply(JSON.stringify);
            resourceInputs["objInt2"] = khulnasoft.output(args ? args.objInt2 : undefined).apply(JSON.stringify);
            resourceInputs["objInt3"] = khulnasoft.output(args ? args.objInt3 : undefined).apply(JSON.stringify);
            resourceInputs["objNum1"] = khulnasoft.output(args ? args.objNum1 : undefined).apply(JSON.stringify);
            resourceInputs["objNum2"] = khulnasoft.output(args ? args.objNum2 : undefined).apply(JSON.stringify);
            resourceInputs["objNum3"] = khulnasoft.output(args ? args.objNum3 : undefined).apply(JSON.stringify);
            resourceInputs["objSecretBool1"] = khulnasoft.output(args?.objSecretBool1 ? khulnasoft.secret(args.objSecretBool1) : undefined).apply(JSON.stringify);
            resourceInputs["objSecretBool2"] = khulnasoft.output(args?.objSecretBool2 ? khulnasoft.secret(args.objSecretBool2) : undefined).apply(JSON.stringify);
            resourceInputs["objSecretBool3"] = khulnasoft.output(args?.objSecretBool3 ? khulnasoft.secret(args.objSecretBool3) : undefined).apply(JSON.stringify);
            resourceInputs["objSecretInt1"] = khulnasoft.output(args?.objSecretInt1 ? khulnasoft.secret(args.objSecretInt1) : undefined).apply(JSON.stringify);
            resourceInputs["objSecretInt2"] = khulnasoft.output(args?.objSecretInt2 ? khulnasoft.secret(args.objSecretInt2) : undefined).apply(JSON.stringify);
            resourceInputs["objSecretInt3"] = khulnasoft.output(args?.objSecretInt3 ? khulnasoft.secret(args.objSecretInt3) : undefined).apply(JSON.stringify);
            resourceInputs["objSecretNum1"] = khulnasoft.output(args?.objSecretNum1 ? khulnasoft.secret(args.objSecretNum1) : undefined).apply(JSON.stringify);
            resourceInputs["objSecretNum2"] = khulnasoft.output(args?.objSecretNum2 ? khulnasoft.secret(args.objSecretNum2) : undefined).apply(JSON.stringify);
            resourceInputs["objSecretNum3"] = khulnasoft.output(args?.objSecretNum3 ? khulnasoft.secret(args.objSecretNum3) : undefined).apply(JSON.stringify);
            resourceInputs["objSecretString1"] = khulnasoft.output(args?.objSecretString1 ? khulnasoft.secret(args.objSecretString1) : undefined).apply(JSON.stringify);
            resourceInputs["objSecretString2"] = khulnasoft.output(args?.objSecretString2 ? khulnasoft.secret(args.objSecretString2) : undefined).apply(JSON.stringify);
            resourceInputs["objSecretString3"] = khulnasoft.output(args?.objSecretString3 ? khulnasoft.secret(args.objSecretString3) : undefined).apply(JSON.stringify);
            resourceInputs["objString1"] = khulnasoft.output(args ? args.objString1 : undefined).apply(JSON.stringify);
            resourceInputs["objString2"] = khulnasoft.output(args ? args.objString2 : undefined).apply(JSON.stringify);
            resourceInputs["objString3"] = khulnasoft.output(args ? args.objString3 : undefined).apply(JSON.stringify);
            resourceInputs["secretBool1"] = khulnasoft.output(args?.secretBool1 ? khulnasoft.secret(args.secretBool1) : undefined).apply(JSON.stringify);
            resourceInputs["secretBool2"] = khulnasoft.output(args?.secretBool2 ? khulnasoft.secret(args.secretBool2) : undefined).apply(JSON.stringify);
            resourceInputs["secretBool3"] = khulnasoft.output(args?.secretBool3 ? khulnasoft.secret(args.secretBool3) : undefined).apply(JSON.stringify);
            resourceInputs["secretInt1"] = khulnasoft.output(args?.secretInt1 ? khulnasoft.secret(args.secretInt1) : undefined).apply(JSON.stringify);
            resourceInputs["secretInt2"] = khulnasoft.output(args?.secretInt2 ? khulnasoft.secret(args.secretInt2) : undefined).apply(JSON.stringify);
            resourceInputs["secretInt3"] = khulnasoft.output(args?.secretInt3 ? khulnasoft.secret(args.secretInt3) : undefined).apply(JSON.stringify);
            resourceInputs["secretNum1"] = khulnasoft.output(args?.secretNum1 ? khulnasoft.secret(args.secretNum1) : undefined).apply(JSON.stringify);
            resourceInputs["secretNum2"] = khulnasoft.output(args?.secretNum2 ? khulnasoft.secret(args.secretNum2) : undefined).apply(JSON.stringify);
            resourceInputs["secretNum3"] = khulnasoft.output(args?.secretNum3 ? khulnasoft.secret(args.secretNum3) : undefined).apply(JSON.stringify);
            resourceInputs["secretString1"] = args?.secretString1 ? khulnasoft.secret(args.secretString1) : undefined;
            resourceInputs["secretString2"] = args?.secretString2 ? khulnasoft.secret(args.secretString2) : undefined;
            resourceInputs["secretString3"] = args?.secretString3 ? khulnasoft.secret(args.secretString3) : undefined;
            resourceInputs["string1"] = args ? args.string1 : undefined;
            resourceInputs["string2"] = args ? args.string2 : undefined;
            resourceInputs["string3"] = args ? args.string3 : undefined;
        }
        opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secretString1", "secretString2", "secretString3"] };
        opts = khulnasoft.mergeOptions(opts, secretOpts);
        super(Provider.__khulnasoftType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    bool1?: khulnasoft.Input<boolean>;
    bool2?: khulnasoft.Input<boolean>;
    bool3?: khulnasoft.Input<boolean>;
    int1?: khulnasoft.Input<number>;
    int2?: khulnasoft.Input<number>;
    int3?: khulnasoft.Input<number>;
    listBool1?: khulnasoft.Input<khulnasoft.Input<boolean>[]>;
    listBool2?: khulnasoft.Input<khulnasoft.Input<boolean>[]>;
    listBool3?: khulnasoft.Input<khulnasoft.Input<boolean>[]>;
    listInt1?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listInt2?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listInt3?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listNum1?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listNum2?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listNum3?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretBool1?: khulnasoft.Input<khulnasoft.Input<boolean>[]>;
    listSecretBool2?: khulnasoft.Input<khulnasoft.Input<boolean>[]>;
    listSecretBool3?: khulnasoft.Input<khulnasoft.Input<boolean>[]>;
    listSecretInt1?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretInt2?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretInt3?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretNum1?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretNum2?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretNum3?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretString1?: khulnasoft.Input<khulnasoft.Input<string>[]>;
    listSecretString2?: khulnasoft.Input<khulnasoft.Input<string>[]>;
    listSecretString3?: khulnasoft.Input<khulnasoft.Input<string>[]>;
    listString1?: khulnasoft.Input<khulnasoft.Input<string>[]>;
    listString2?: khulnasoft.Input<khulnasoft.Input<string>[]>;
    listString3?: khulnasoft.Input<khulnasoft.Input<string>[]>;
    mapBool1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<boolean>}>;
    mapBool2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<boolean>}>;
    mapBool3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<boolean>}>;
    mapInt1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapInt2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapInt3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapNum1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapNum2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapNum3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretBool1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<boolean>}>;
    mapSecretBool2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<boolean>}>;
    mapSecretBool3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<boolean>}>;
    mapSecretInt1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretInt2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretInt3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretNum1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretNum2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretNum3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretString1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<string>}>;
    mapSecretString2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<string>}>;
    mapSecretString3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<string>}>;
    mapString1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<string>}>;
    mapString2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<string>}>;
    mapString3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<string>}>;
    num1?: khulnasoft.Input<number>;
    num2?: khulnasoft.Input<number>;
    num3?: khulnasoft.Input<number>;
    objBool1?: khulnasoft.Input<inputs.Tbool1Args>;
    objBool2?: khulnasoft.Input<inputs.Tbool2Args>;
    objBool3?: khulnasoft.Input<inputs.Tbool3Args>;
    objInt1?: khulnasoft.Input<inputs.Tint1Args>;
    objInt2?: khulnasoft.Input<inputs.Tint2Args>;
    objInt3?: khulnasoft.Input<inputs.Tint3Args>;
    objNum1?: khulnasoft.Input<inputs.Tnum1Args>;
    objNum2?: khulnasoft.Input<inputs.Tnum2Args>;
    objNum3?: khulnasoft.Input<inputs.Tnum3Args>;
    objSecretBool1?: khulnasoft.Input<inputs.TsecretBool1Args>;
    objSecretBool2?: khulnasoft.Input<inputs.TsecretBool2Args>;
    objSecretBool3?: khulnasoft.Input<inputs.TsecretBool3Args>;
    objSecretInt1?: khulnasoft.Input<inputs.TsecretInt1Args>;
    objSecretInt2?: khulnasoft.Input<inputs.TsecretInt2Args>;
    objSecretInt3?: khulnasoft.Input<inputs.TsecretInt3Args>;
    objSecretNum1?: khulnasoft.Input<inputs.TsecretNum1Args>;
    objSecretNum2?: khulnasoft.Input<inputs.TsecretNum2Args>;
    objSecretNum3?: khulnasoft.Input<inputs.TsecretNum3Args>;
    objSecretString1?: khulnasoft.Input<inputs.TsecretString1Args>;
    objSecretString2?: khulnasoft.Input<inputs.TsecretString2Args>;
    objSecretString3?: khulnasoft.Input<inputs.TsecretString3Args>;
    objString1?: khulnasoft.Input<inputs.Tstring1Args>;
    objString2?: khulnasoft.Input<inputs.Tstring2Args>;
    objString3?: khulnasoft.Input<inputs.Tstring3Args>;
    secretBool1?: khulnasoft.Input<boolean>;
    secretBool2?: khulnasoft.Input<boolean>;
    secretBool3?: khulnasoft.Input<boolean>;
    secretInt1?: khulnasoft.Input<number>;
    secretInt2?: khulnasoft.Input<number>;
    secretInt3?: khulnasoft.Input<number>;
    secretNum1?: khulnasoft.Input<number>;
    secretNum2?: khulnasoft.Input<number>;
    secretNum3?: khulnasoft.Input<number>;
    secretString1?: khulnasoft.Input<string>;
    secretString2?: khulnasoft.Input<string>;
    secretString3?: khulnasoft.Input<string>;
    string1?: khulnasoft.Input<string>;
    string2?: khulnasoft.Input<string>;
    string3?: khulnasoft.Input<string>;
}
