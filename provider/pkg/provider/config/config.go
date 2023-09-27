package config

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-xyz/provider/pkg/provider/api"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type XyzConfig struct {
	Url    string `pulumi:"vcoUrl"`
	ApiKey string `pulumi:"vcoApiKey" provider:"secret"`

	client api.Client
}

var _ = (infer.CustomCheck[*XyzConfig])((*XyzConfig)(nil))

func (c *XyzConfig) Check(ctx p.Context, name string, oldInputs, newInputs resource.PropertyMap) (*XyzConfig, []p.CheckFailure, error) {
	c.Url = newInputs["vcoUrl"].StringValue()
	c.ApiKey = newInputs["vcoApiKey"].StringValue()

	return c, []p.CheckFailure{}, nil
}

var _ = (infer.Annotated)((*XyzConfig)(nil))

func (c *XyzConfig) Annotate(a infer.Annotator) {
	a.Describe(&c.Url, "FQDN of the VCO")
	a.Describe(&c.ApiKey, "API key for the VCO")
}

var _ = (infer.CustomConfigure)((*XyzConfig)(nil))

func (c *XyzConfig) Configure(ctx p.Context) error {
	c.client = api.NewClient(c.Url, c.ApiKey, 10)
	return nil
}

func (c *XyzConfig) Client() *api.Client {
	return &c.client
}
