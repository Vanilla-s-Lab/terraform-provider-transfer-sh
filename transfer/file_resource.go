package transfer

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"os"
	"path"

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

			// https://github.com/dutchcoders/transfer.sh
			"link": schema.StringAttribute{Computed: true},
		},
	}
}

const URL = "https://transfer.sh"

func (f fileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan fileResource
	req.Plan.Get(ctx, &plan)

	filePath := plan.Path.ValueString()

	// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		resp.Diagnostics.AddError("File not found", filePath)
		return
	}

	fileName := path.Base(filePath)
	url := URL + "/" + fileName

	buf, _ := os.Open(filePath)

	// https://pkg.go.dev/crypto/md5#example-New-File
	h := md5.New()
	_, _ = io.Copy(h, buf)

	//https://stackoverflow.com/questions/2377881/how-to-get-a-md5-hash-from-a-string-in-golang
	hash := h.Sum(nil)
	hashString := hex.EncodeToString(hash[:])

	plan.Hash = types.StringValue(hashString)

	_, _ = buf.Seek(0, 0)

	// https://pkg.go.dev/net/http

	request, _ := http.NewRequest(http.MethodPut, url, buf)
	response, _ := (&http.Client{}).Do(request)

	body, _ := io.ReadAll(response.Body)
	plan.Link = types.StringValue(string(body))

	resp.State.Set(ctx, plan)
}

func (f fileResource) Read(context.Context, resource.ReadRequest, *resource.ReadResponse) {
}

func (f fileResource) Update(context.Context, resource.UpdateRequest, *resource.UpdateResponse) {
}

func (f fileResource) Delete(context.Context, resource.DeleteRequest, *resource.DeleteResponse) {
}
