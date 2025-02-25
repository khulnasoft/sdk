// Code generated by khulnasoft-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"example.com/khulnasoft-config/sdk/go/v9/config/internal"
	"github.com/khulnasoft/sdk/go/khulnasoft"
	"github.com/khulnasoft/sdk/go/khulnasoft/config"
)

var _ = internal.GetEnvOrDefault

func GetName(ctx *khulnasoft.Context) string {
	return config.Get(ctx, "config:name")
}
func GetPluginDownloadURL(ctx *khulnasoft.Context) string {
	return config.Get(ctx, "config:pluginDownloadURL")
}
