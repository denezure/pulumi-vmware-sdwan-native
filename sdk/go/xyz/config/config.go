// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"internal"
)

var _ = internal.GetEnvOrDefault

// API key for the VCO
func GetVcoApiKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "xyz:vcoApiKey")
}

// FQDN of the VCO
func GetVcoUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "xyz:vcoUrl")
}
