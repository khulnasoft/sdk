package main

import (
	"example.com/khulnasoft-simple-invoke/sdk/go/v10/simpleinvoke"
	"github.com/khulnasoft/sdk/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := simpleinvoke.NewProvider(ctx, "explicitProvider", nil)
		if err != nil {
			return err
		}
		first, err := simpleinvoke.NewStringResource(ctx, "first", &simpleinvoke.StringResourceArgs{
			Text: khulnasoft.String("first hello"),
		})
		if err != nil {
			return err
		}
		data := simpleinvoke.MyInvokeOutput(ctx, simpleinvoke.MyInvokeOutputArgs{
			Value: khulnasoft.String("hello"),
		}, khulnasoft.DependsOn([]khulnasoft.Resource{
			khulnasoft.Resource(first),
		}))
		_, err = simpleinvoke.NewStringResource(ctx, "second", &simpleinvoke.StringResourceArgs{
			Text: data.ApplyT(func(data simpleinvoke.MyInvokeResult) (string, error) {
				return data.Result, nil
			}).(khulnasoft.StringOutput),
		})
		if err != nil {
			return err
		}
		ctx.Export("hello", data.ApplyT(func(data simpleinvoke.MyInvokeResult) (string, error) {
			return data.Result, nil
		}).(khulnasoft.StringOutput))
		return nil
	})
}
