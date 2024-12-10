package main

import (
	"github.com/khulnasoft/sdk/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		ctx.Export("output_true", khulnasoft.Bool(true))
		return nil
	})
}
