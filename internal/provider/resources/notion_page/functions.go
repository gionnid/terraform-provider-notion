package notion_page

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (r *NotionPage) EvaluateState(response *http.Response, ctx context.Context) (state NotionPageResourceModel, archived bool, err error) {

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

	// Extract ID
	if id, ok := responseData["id"].(string); ok {
		state.ID = types.StringValue(id)
	}

	// Extract parent ID
	if parent, ok := responseData["parent"].(map[string]interface{}); ok {
		if pageID, ok := parent["page_id"].(string); ok {
			state.ParentID = types.StringValue(pageID)
		}
	}

	// Get arc_val value
	archived = false
	if arc_val, ok := responseData["archived"].(bool); ok {
		archived = arc_val
	}

	tflog.Debug(ctx, "State: Name -> "+state.Name.ValueString()+" ID -> "+state.ID.ValueString()+" ParentID -> "+state.ParentID.ValueString())
	return state, archived, nil
}
