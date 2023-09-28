package provider

import (
	"strings"

	"github.com/blang/semver"
	"github.com/nick-barrett/pulumi-veco/provider/pkg/provider/config"
	resources "github.com/nick-barrett/pulumi-veco/provider/pkg/provider/resources"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
)

const Name string = "veco"

func Provider() p.Provider {
	prov := infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			DisplayName: "veco",
			Description: "VMware Edge Cloud Orchestrator provider",
			Keywords: []string{
				"pulumi", "vmware", "sd-wan", "kind/native",
			},
			LanguageMap: map[string]any{
				"nodejs": map[string]any{
					"dependencies": map[string]string{
						"@pulumi/pulumi": "^3.0.0",
					},
				},
				"go": map[string]any{},
			},
		},
		Resources: []infer.InferredResource{
			infer.Resource[*resources.AddressGroup, resources.AddressGroupInputs, resources.AddressGroupOutputs](),
			infer.Resource[*resources.ServiceGroup, resources.ServiceGroupInputs, resources.ServiceGroupOutputs](),
		},
		Config: infer.Config[*config.VecoConfig](),
	})

	return prov
}

func Schema(version string) (string, error) {
	version = strings.TrimPrefix(version, "v")
	s, err := integration.NewServer(Name, semver.MustParse(version), Provider()).
		GetSchema(p.GetSchemaRequest{})
	return s.Schema, err
}
