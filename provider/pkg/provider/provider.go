package provider

import (
	"strings"

	"github.com/blang/semver"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi-xyz/provider/pkg/provider/config"
	resources "github.com/pulumi/pulumi-xyz/provider/pkg/provider/resources"
)

const Name string = "xyz"

func Provider() p.Provider {
	prov := infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			DisplayName: "xyz",
			Description: "Work-in-progress VMware SD-WAN Pulumi provider",
			Keywords: []string{
				"pulumi", "vmware", "sd-wan", "kind/native",
			},
			LanguageMap: map[string]any{
				"nodejs": map[string]any{
					"dependencies": map[string]string{
						"@pulumi/pulumi": "^3.0.0",
					},
				},
			},
		},
		Resources: []infer.InferredResource{
			infer.Resource[*resources.AddressGroup, resources.AddressGroupInputs, resources.AddressGroupOutputs](),
			infer.Resource[*resources.ServiceGroup, resources.ServiceGroupInputs, resources.ServiceGroupOutputs](),
		},
		Config: infer.Config[*config.XyzConfig](),
	})

	return prov
}

func Schema(version string) (string, error) {
	version = strings.TrimPrefix(version, "v")
	s, err := integration.NewServer(Name, semver.MustParse(version), Provider()).
		GetSchema(p.GetSchemaRequest{})
	return s.Schema, err
}
