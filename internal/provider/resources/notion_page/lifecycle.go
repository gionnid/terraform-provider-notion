package notion_page

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func (r *NotionPage) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Starting Create operation for Notion Page")

	var plan, state NotionPageResourceModel
	resp.Diagnostics.Append(
		req.Plan.Get(ctx, &plan)...,
	)
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

	if !r.NotionApiClient.HandleApiResponse(apiResponse, err, "Failed to create page", resp.Diagnostics.AddError) {
		return
	}

	state, _, err = r.EvaluateState(apiResponse, ctx)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get page state", err.Error())
		return
	}

	// modelsManager.SetState(&resp.Diagnostics, &resp.State, ctx, &state)
	resp.Diagnostics.Append(
		resp.State.Set(ctx, &state)...,
	)
	tflog.Debug(ctx, "Completed Create operation for Notion Page")
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *NotionPage) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Starting Read operation for Notion Page")

	// Obtain the current state of the resource
	var state NotionPageResourceModel
	resp.Diagnostics.Append(
		req.State.Get(ctx, &state)...,
	)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read the resource from Notion
	url := "https://api.notion.com/v1/pages/" + state.ID.ValueString()
	tflog.Debug(ctx, "Read URL: "+url)
	apiResponse, err := r.NotionApiClient.Get(url)
	if !r.NotionApiClient.HandleApiResponse(apiResponse, err, "Failed to read page", resp.Diagnostics.AddError) {
		return
	}

	// Update the resource state
	newState, archived, err := r.EvaluateState(apiResponse, ctx)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get page state", err.Error())
		return
	}
	if archived {
		resp.State.RemoveResource(ctx)
		return
	}
	resp.Diagnostics.Append(
		resp.State.Set(ctx, &newState)...,
	)
}

func (r *NotionPage) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Starting Update operation for Notion Page")

	// Get planned changes
	var plan, state NotionPageResourceModel
	resp.Diagnostics.Append(
		req.Plan.Get(ctx, &plan)...,
	)
	resp.Diagnostics.Append(
		req.State.Get(ctx, &state)...,
	)
	if plan.ParentID.ValueString() != state.ParentID.ValueString() {
		resp.Diagnostics.AddError(
			"Parent ID Cannot Be Changed",
			"The parent_id of a Notion page cannot be changed after creation. Current parent_id: "+state.ParentID.ValueString()+", Requested parent_id: "+plan.ParentID.ValueString(),
		)
	}
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
	if !r.NotionApiClient.HandleApiResponse(apiResponse, err, "Failed to update page", resp.Diagnostics.AddError) {
		return
	}

	state, _, err = r.EvaluateState(apiResponse, ctx)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get page state", err.Error())
		return
	}

	resp.Diagnostics.Append(
		resp.State.Set(ctx, &state)...,
	)
}

func (r *NotionPage) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Starting Delete operation for Notion Page")

	var state NotionPageResourceModel
	resp.Diagnostics.Append(
		req.State.Get(ctx, &state)...,
	)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Delete ID: "+state.ID.ValueString())
	body := `{"archived": true}`
	apiResponse, err := r.NotionApiClient.Patch(
		"https://api.notion.com/v1/pages/"+state.ID.ValueString(),
		body,
	)

	if !r.NotionApiClient.HandleApiResponse(apiResponse, err, "Failed to delete page", resp.Diagnostics.AddError) {
		return
	}
}
