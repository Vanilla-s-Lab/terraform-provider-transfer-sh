package transfer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func New() provider.Provider {
	return &transferProvider{}
}

var (
	_ provider.Provider = &transferProvider{}
)

type transferProvider struct{}

func (t transferProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "transfer-sh"
}

func (t transferProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (t transferProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}

func (t transferProvider) DataSources(context.Context) []func() datasource.DataSource {
	return nil
}

func (t transferProvider) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewFileResource,
	}
}
