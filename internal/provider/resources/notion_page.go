package provider_resources

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gionnid/terraform-provider-notion/internal/provider/client"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &NotionPage{}

type NotionPage struct {
	NotionApiClient *client.NotionApiClient

	ID       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	ParentID types.String `tfsdk:"parent_id"`
}

type NotionPageResourceModel struct {
	ID       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	ParentID types.String `tfsdk:"parent_id"`
}

func NewNotionPage(notion_client *client.NotionApiClient) resource.Resource {
	return &NotionPage{
		NotionApiClient: notion_client,
	}
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

	body := `{
    "parent": {"page_id": "` + plan.ParentID.ValueString() + `"},
    "properties":{"title":{"title":[{"text":{"content":"` + plan.Name.ValueString() + `"}}]}}
	}`
	apiResponse, err := r.NotionApiClient.Post(
		"https://api.notion.com/v1/pages",
		body,
	)

	if err != nil || apiResponse.StatusCode != 200 {
		resp.Diagnostics.AddError("Failed to create page", err.Error())
		return
	}

	state, err := r.GetState(apiResponse, ctx, resp)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get page state", err.Error())
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	tflog.Warn(ctx, "Completed Create operation for Notion Page")
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Body: "+body)
	tflog.Info(ctx, "Successfully set resource state")
}

func (r *NotionPage) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Implement resource read logic
}

func (r *NotionPage) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Implement resource update logic
}

func (r *NotionPage) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Starting Delete operation for Notion Page")

	var state NotionPageResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Delete ID: "+state.ID.ValueString())

	body := `{"archived": true}`
	apiResponse, err := r.NotionApiClient.Patch(
		"https://api.notion.com/v1/pages/"+state.ID.ValueString(),
		body,
	)

	if err != nil || apiResponse.StatusCode != 200 {
		resp.Diagnostics.AddError("Failed to delete page", err.Error())
		return
	}

	tflog.Info(ctx, "Successful Delete")
}

func (r *NotionPage) GetState(response *http.Response, ctx context.Context, resp *resource.CreateResponse) (state NotionPageResourceModel, err error) {
	if response.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(response.Body)
		resp.Diagnostics.AddError("Failed to get page details", fmt.Sprintf("status code %d: %s", response.StatusCode, string(bodyBytes)))
		return state, fmt.Errorf("status code %d", response.StatusCode)
	}

	var responseData map[string]interface{}
	json.NewDecoder(response.Body).Decode(&responseData)
	defer response.Body.Close()

	// Extract title from properties
	if properties, ok := responseData["properties"].(map[string]interface{}); ok {
		if titleProp, ok := properties["title"].(map[string]interface{}); ok {
			if titleArr, ok := titleProp["title"].([]interface{}); ok && len(titleArr) > 0 {
				if titleObj, ok := titleArr[0].(map[string]interface{}); ok {
					if textObj, ok := titleObj["text"].(map[string]interface{}); ok {
						if content, ok := textObj["content"].(string); ok {
							state.Name = types.StringValue(content)
						}
					}
				}
			}
		}
	}

	if id, ok := responseData["id"].(string); ok {
		state.ID = types.StringValue(id)
	}

	// Extract parent ID
	if parent, ok := responseData["parent"].(map[string]interface{}); ok {
		if pageID, ok := parent["page_id"].(string); ok {
			state.ParentID = types.StringValue(pageID)
		}
	}
	tflog.Debug(ctx, "State: Name -> "+state.Name.ValueString()+" ID -> "+state.ID.ValueString()+" ParentID -> "+state.ParentID.ValueString())
	return state, nil
}
