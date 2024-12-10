package main

import (
	"github.com/khulnasoft/sdk/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		ctx.Export("stackOutput", khulnasoft.String(ctx.Stack()))
		ctx.Export("projectOutput", khulnasoft.String(ctx.Project()))
		ctx.Export("organizationOutput", khulnasoft.String(ctx.Organization()))
		return nil
	})
}
