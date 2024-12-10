package main

import (
	"example.com/khulnasoft-secret/sdk/go/v14/secret"
	"github.com/khulnasoft/sdk/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := secret.NewResource(ctx, "res", &secret.ResourceArgs{
			Private: khulnasoft.String("closed"),
			Public:  khulnasoft.String("open"),
			PrivateData: &secret.DataArgs{
				Private: khulnasoft.String("closed"),
				Public:  khulnasoft.String("open"),
			},
			PublicData: &secret.DataArgs{
				Private: khulnasoft.String("closed"),
				Public:  khulnasoft.String("open"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
