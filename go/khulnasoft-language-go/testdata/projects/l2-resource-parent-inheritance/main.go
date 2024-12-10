package main

import (
	"example.com/khulnasoft-simple/sdk/go/v2/simple"
	"github.com/khulnasoft/sdk/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		provider, err := simple.NewProvider(ctx, "provider", nil)
		if err != nil {
			return err
		}
		parent1, err := simple.NewResource(ctx, "parent1", &simple.ResourceArgs{
			Value: khulnasoft.Bool(true),
		}, khulnasoft.Provider(provider))
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "child1", &simple.ResourceArgs{
			Value: khulnasoft.Bool(true),
		}, khulnasoft.Parent(parent1))
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "orphan1", &simple.ResourceArgs{
			Value: khulnasoft.Bool(true),
		})
		if err != nil {
			return err
		}
		parent2, err := simple.NewResource(ctx, "parent2", &simple.ResourceArgs{
			Value: khulnasoft.Bool(true),
		}, khulnasoft.Protect(true))
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "child2", &simple.ResourceArgs{
			Value: khulnasoft.Bool(true),
		}, khulnasoft.Parent(parent2))
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "orphan2", &simple.ResourceArgs{
			Value: khulnasoft.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
