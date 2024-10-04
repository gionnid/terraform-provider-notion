package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.Resource = &NotionResource{}

type NotionResource struct {
	// Add any necessary resource-level fields here
}

func NewNotionResource() resource.Resource {
	return &NotionResource{}
}

func (r *NotionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_first_resource"
}

func (r *NotionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *NotionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Implement resource creation logic
}

func (r *NotionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Implement resource read logic
}

func (r *NotionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Implement resource update logic
}

func (r *NotionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Implement resource deletion logic
}
