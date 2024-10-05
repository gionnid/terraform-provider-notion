package tests

import (
	"context"
	"testing"

	"github.com/gionnid/terraform-provider-notion/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func TestProviderSchema(t *testing.T) {
	server := providerserver.NewProtocol6(provider.New())()
	schema, _ := server.GetProviderSchema(context.Background(), &tfprotov6.GetProviderSchemaRequest{})
	if !AttributeInProviderSchema(schema, "notion_integration_token") {
		t.Error("Expected 'notion_integration_token' attribute in provider schema, but it was not found")
	}

	if !AttributeInProviderSchema(schema, "notion_api_version") {
		t.Error("Expected 'notion_api_version' attribute in provider schema, but it was not found")
	}
}

func AttributeInProviderSchema(schema *tfprotov6.GetProviderSchemaResponse, name string) bool {
	for _, attr := range schema.Provider.Block.Attributes {
		if attr.Name == name {
			return true
		}
	}
	return false
}
