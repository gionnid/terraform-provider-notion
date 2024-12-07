package client_test

import (
	"context"
	"testing"

	"github.com/gionnid/terraform-provider-notion/internal/provider/client"
	"github.com/stretchr/testify/assert"
)

func TestNewNotionApiClient(t *testing.T) {
	version := "2022-06-28"
	token := "secret_1234567890"
	client.NewNotionApiClient(context.Background(), version, token)
	assert.NotNil(t, client.GetNotionApiClient())
	assert.Equal(t, version, client.GetNotionApiClient().NotionApiVersion)
}

func TestPostSuccessful(t *testing.T) {
	notion_client := client.NewNotionApiClient(context.Background(), "2022-06-28", "secret_1234567890")

	headers := notion_client.GetHeaders(true)

	body := map[string]interface{}{
		"name": "John Doe",
	}

	response, err := notion_client.Post("http://0.0.0.0:8000", headers, body)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestPostFails(t *testing.T) {
	notion_client := client.NewNotionApiClient(context.Background(), "2099-99-99", "secret_1234567890")

	headers := notion_client.GetHeaders(true)

	body := map[string]interface{}{
		"name": "John Doe",
	}

	response, err := notion_client.Post("http://invalid-url.com", headers, body)

	assert.Error(t, err)
	assert.Nil(t, response)
}
