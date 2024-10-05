package resources_test

import (
	"testing"

	provider_resources "github.com/gionnid/terraform-provider-notion/internal/provider/resources"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestNotionPage(t *testing.T) {
	// Create a new NotionPage
	page := provider_resources.NotionPage{
		ID:       types.StringValue("page-id-123"),
		Name:     types.StringValue("Test Page"),
		ParentID: types.StringValue("parent-id-123"),
	}

	assert.Equal(t, "page-id-123", page.ID, "Page ID should match")
	assert.Equal(t, "Test Page", page.Name, "Page Title should match")
	assert.Equal(t, "parent-id-123", page.ParentID, "Page URL should match")
}
