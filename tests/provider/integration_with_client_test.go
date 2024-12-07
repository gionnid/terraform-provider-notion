package tests

import (
	"context"
	"testing"

	"github.com/gionnid/terraform-provider-notion/internal/provider"
	"github.com/gionnid/terraform-provider-notion/internal/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func TestUninitializedClient(t *testing.T) {
	notion_client := client.NewNotionApiClient(context.Background())
	server := providerserver.NewProtocol6(provider.New(notion_client))()
	schema, _ := server.GetProviderSchema(context.Background(), &tfprotov6.GetProviderSchemaRequest{})
	if !AttributeInProviderSchema(schema, "notion_integration_token") {
		t.Error("Expected 'notion_integration_token' attribute in provider schema, but it was not found")
	}

	if !AttributeInProviderSchema(schema, "notion_api_version") {
		t.Error("Expected 'notion_api_version' attribute in provider schema, but it was not found")
	}
}
