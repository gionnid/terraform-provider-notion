package provider_resources

import (
	"github.com/gionnid/terraform-provider-notion/internal/provider/client"
	"github.com/gionnid/terraform-provider-notion/internal/provider/resources/notion_page"
)

func NewNotionPage(client *client.NotionApiClient) *notion_page.NotionPage {
	return &notion_page.NotionPage{NotionApiClient: client}
}
