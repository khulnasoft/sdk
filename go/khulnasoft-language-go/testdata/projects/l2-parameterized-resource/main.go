package main

import (
	"example.com/khulnasoft-subpackage/sdk/go/v2/subpackage"
	"github.com/khulnasoft/sdk/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		// The resource name is based on the parameter value
		example, err := subpackage.NewHelloWorld(ctx, "example", nil)
		if err != nil {
			return err
		}
		ctx.Export("parameterValue", example.ParameterValue)
		return nil
	})
}
