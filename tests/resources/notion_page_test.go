package resources_tests

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	provider_resources "github.com/gionnid/terraform-provider-notion/internal/provider/resources"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

var (
	title     = "Test Page"
	parent_id = "parent-id-123"
	id        = "page-id-123"
)

func TestNotionPage(t *testing.T) {
	// Create a new NotionPage
	page := provider_resources.NotionPage{
		ID:       types.StringValue(id),
		Name:     types.StringValue(title),
		ParentID: types.StringValue(parent_id),
	}

	assert.Equal(t, id, page.ID.ValueString(), "Page ID should match")
	assert.Equal(t, title, page.Name.ValueString(), "Page Title should match")
	assert.Equal(t, parent_id, page.ParentID.ValueString(), "Page URL should match")
}

func TestGetPageState(t *testing.T) {

	response := &http.Response{
		StatusCode: 200,
		Body: io.NopCloser(strings.NewReader(`{
    "parent": {"page_id": "` + parent_id + `"},
	"id": "` + id + `",
    "properties":{"title":{"title":[{"text":{"content":"` + title + `"}}]}
	}}`)),
	}

	page := &provider_resources.NotionPage{}

	state, err := page.GetState(response, context.Background(), nil)
	assert.Nil(t, err)
	assert.Equal(t, id, state.ID.ValueString())
	assert.Equal(t, title, state.Name.ValueString())
	assert.Equal(t, parent_id, state.ParentID.ValueString())
}
