package main

import (
	"github.com/khulnasoft/sdk/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		ctx.Export("output_true", khulnasoft.Bool(true))
		ctx.Export("output_false", khulnasoft.Bool(false))
		return nil
	})
}
