package transfer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	_ resource.Resource = &fileResource{}
)

func NewFileResource() resource.Resource {
	return &fileResource{}
}

type fileResource struct{}

func (f fileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file"
}

func (f fileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (f fileResource) Create(context.Context, resource.CreateRequest, *resource.CreateResponse) {
}

func (f fileResource) Read(context.Context, resource.ReadRequest, *resource.ReadResponse) {
}

func (f fileResource) Update(context.Context, resource.UpdateRequest, *resource.UpdateResponse) {
}

func (f fileResource) Delete(context.Context, resource.DeleteRequest, *resource.DeleteResponse) {
}
