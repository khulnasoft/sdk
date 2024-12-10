//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/sdk/go/khulnasoft"
	"github.com/khulnasoft/sdk/go/khulnasoft/config"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		c := config.New(ctx, "")
		ctx.Export("exp_static", khulnasoft.String("foo"))
		ctx.Export("exp_cfg", khulnasoft.String(c.Get("bar")))
		ctx.Export("exp_secret", c.GetSecret("buzz"))
		return nil
	})
}
