package notion_page

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

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

	state, _, err := r.GetState(apiResponse, ctx)
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
	tflog.Debug(ctx, "Starting Read operation for Notion Page")

	// Obtain the current state of the resource
	var state NotionPageResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError("Failed to read state", diags.Errors()[0].Summary())
		return
	}

	// Read the resource from Notion
	url := "https://api.notion.com/v1/pages/" + state.ID.ValueString()
	tflog.Debug(ctx, "Read URL: "+url)
	apiResponse, err := r.NotionApiClient.Get(url)
	if !r.HandleApiResponse(apiResponse, err, "Failed to read page", resp.Diagnostics.AddError) {
		return
	}

	// Update the resource state
	newState, archived, err := r.GetState(apiResponse, ctx)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get page state", err.Error())
		return
	}
	tflog.Debug(ctx, "Archived: "+strconv.FormatBool(archived))
	if archived {
		resp.State.RemoveResource(ctx)
		return
	}
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
}

func (r *NotionPage) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Starting Update operation for Notion Page")

	// Get planned changes
	var plan NotionPageResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Update planned:")
	tflog.Info(ctx, "\t plan ID: "+plan.ID.ValueString())
	tflog.Info(ctx, "\t plan NAME: "+plan.Name.ValueString())
	tflog.Info(ctx, "\t plan PARENT ID: "+plan.ParentID.ValueString())

	body := `{
	    "properties": {
	        "title": {
	            "title": [
	                {
	                    "text": {
	                        "content": "` + plan.Name.ValueString() + `"
	                    }
	                }
	            ]
	        }
	    }
	}`
	url := "https://api.notion.com/v1/pages/" + plan.ID.ValueString()
	tflog.Debug(ctx, "Update URL: "+url)
	apiResponse, err := r.NotionApiClient.Patch(url, body)
	if !r.HandleApiResponse(apiResponse, err, "Failed to update page", resp.Diagnostics.AddError) {
		return
	}

	state, _, err := r.GetState(apiResponse, ctx)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get page state", err.Error())
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
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

	if err != nil {
		resp.Diagnostics.AddError("Failed to delete page", err.Error())
		return
	}
	if apiResponse.StatusCode != 200 {
		resp.Diagnostics.AddError("Failed to delete page", "Status Code: "+strconv.Itoa(apiResponse.StatusCode))
		return
	}

	tflog.Info(ctx, "Successful Delete")
}
