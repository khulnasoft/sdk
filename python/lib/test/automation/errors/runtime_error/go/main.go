//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/sdk/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		var x []string
		ctx.Export("a", khulnasoft.String(x[0]))
		return nil
	})
}
