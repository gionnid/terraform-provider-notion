package notion_page

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gionnid/terraform-provider-notion/internal/provider/resources/notion_page"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)



func TestNotionPage(t *testing.T) {
	// Create a new NotionPage
	page := notion_page.NotionPage{
		ID:       types.StringValue(id),
		Name:     types.StringValue(title),
		ParentID: types.StringValue(parent_id),
	}

	assert.Equal(t, id, page.ID.ValueString(), "Page ID should match")
	assert.Equal(t, title, page.Name.ValueString(), "Page Title should match")
	assert.Equal(t, parent_id, page.ParentID.ValueString(), "Page URL should match")
}

func TestGetPageState(t *testing.T) {

	apiResponse := &http.Response{
		StatusCode: 200,
		Body: io.NopCloser(strings.NewReader(`{
    "parent": {"page_id": "` + parent_id + `"},
	"id": "` + id + `",
    "properties":{"title":{"title":[{"text":{"content":"` + title + `"}}]}
	}}`)),
	}

	page := &notion_page.NotionPage{}

	state, _, err := page.GetState(apiResponse, context.Background())
	assert.Nil(t, err)
	assert.Equal(t, id, state.ID.ValueString())
	assert.Equal(t, title, state.Name.ValueString())
	assert.Equal(t, parent_id, state.ParentID.ValueString())
}

func TestGetPageWithCreateResponse(t *testing.T) {
	apiResponse := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{}`)),
	}

	page := &notion_page.NotionPage{}
	state, _, err := page.GetState(apiResponse, context.Background())

	assert.Nil(t, err)
	assert.NotNil(t, state)
}
