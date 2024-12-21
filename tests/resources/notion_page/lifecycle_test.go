package notion_page

import (
	"context"
	"testing"

	"io"
	"net/http"
	"strings"

	"github.com/gionnid/terraform-provider-notion/internal/provider/resources/notion_page"
	"github.com/gionnid/terraform-provider-notion/tests/mocks"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func MockedPost() (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body: io.NopCloser(strings.NewReader(`{
			"id": "` + id + `",
			"properties": {
				"title": {
					"title": [
						{
							"text": {
								"content": "` + title + `"
							}
						}
					]
				}
			},
			"parent": {
				"page_id": "` + parent_id + `"
			},
			"archived": false
		}`)),
	}, nil
}

func TestNotionPage_Simple(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNotionClient := mocks.NewMockNotionAPI(ctrl)
	mockNotionClient.EXPECT().Post(gomock.Any(), gomock.Any()).Return(MockedPost())

	page := notion_page.NotionPage{NotionApiClient: mockNotionClient}

	testPlan := tfsdk.Plan{
		Raw: tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"parent_id": tftypes.String,
				"id":        tftypes.String,
				"name":      tftypes.String,
			},
		}, map[string]tftypes.Value{
			"parent_id": tftypes.NewValue(tftypes.String, parent_id),
			"id":        tftypes.NewValue(tftypes.String, id),
			"name":      tftypes.NewValue(tftypes.String, title),
		}),
		Schema: page.GetSchema(),
	}

	ctx := context.Background()
	resp := resource.CreateResponse{
		State: tfsdk.State{
			Schema: page.GetSchema(),
		},
	}
	req := resource.CreateRequest{
		Plan: testPlan,
	}

	page.Create(ctx, req, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Create operation failed: %v", resp.Diagnostics)
	}

	var statePost notion_page.NotionPageResourceModel
	resp.State.Get(ctx, &statePost)
	assert.Equal(t, statePost.ID.ValueString(), id)
	assert.Equal(t, statePost.Name.ValueString(), title)
	assert.Equal(t, statePost.ParentID.ValueString(), parent_id)
}
