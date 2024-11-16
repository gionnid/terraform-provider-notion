package provider_resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &NotionPage{}

type NotionPage struct {
	ID       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	ParentID types.String `tfsdk:"parent_id"`
}

type NotionPageResourceModel struct {
	ID       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	ParentID types.String `tfsdk:"parent_id"`
}

func NewNotionPage() resource.Resource {
	return &NotionPage{}
}

func (r *NotionPage) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_page"
}

func (r *NotionPage) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"parent_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *NotionPage) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create operation for Notion Page")

	// Retrieve the plan data
	var plan NotionPageResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	name := plan.Name.ValueString()
	parent_id := plan.ParentID.ValueString()

	// Create a new state to hold the resource data
	var state NotionPageResourceModel

	// Set the name in the state
	// Note: In a real implementation, you'd get this from the API response
	state.Name = types.StringValue(name)

	// Set a placeholder ID
	// Note: In a real implementation, you'd get this from the API response
	state.ID = types.StringValue("page_id_from_api")
	state.ParentID = types.StringValue(parent_id)

	// Save the state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	tflog.Warn(ctx, "Completed Create operation for Notion Page")
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Successfully set resource state")
}

func (r *NotionPage) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Implement resource read logic
}

func (r *NotionPage) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Implement resource update logic
}

func (r *NotionPage) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Implement resource deletion logic
}
