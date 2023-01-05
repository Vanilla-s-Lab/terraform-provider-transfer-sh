package transfer

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	_ resource.Resource = &fileResource{}
)

func NewFileResource() resource.Resource {
	return &fileResource{}
}

type fileResource struct {
	Path types.String `tfsdk:"file_path"`
	Hash types.String `tfsdk:"file_hash"`
	Link types.String `tfsdk:"link"`
}

func (f fileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file"
}

func (f fileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{

			// https://registry.terraform.io/providers/linode/linode/latest/docs/resources/image
			"file_path": schema.StringAttribute{Required: true},
			"file_hash": schema.StringAttribute{Computed: true},

			//https://github.com/dutchcoders/transfer.sh
			"link": schema.StringAttribute{Computed: true},
		},
	}
}

func (f fileResource) Create(context.Context, resource.CreateRequest, *resource.CreateResponse) {
}

func (f fileResource) Read(context.Context, resource.ReadRequest, *resource.ReadResponse) {
}

func (f fileResource) Update(context.Context, resource.UpdateRequest, *resource.UpdateResponse) {
}

func (f fileResource) Delete(context.Context, resource.DeleteRequest, *resource.DeleteResponse) {
}
