//go:build !all
// +build !all

package main

import (
	"time"

	"github.com/khulnasoft/sdk/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		time.Sleep(1 * time.Second)
		ctx.Export("exp_static", khulnasoft.String("foo"))
		return nil
	})
}
