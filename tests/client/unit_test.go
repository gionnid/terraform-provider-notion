package client_test

import (
	"testing"

	"github.com/gionnid/terraform-provider-notion/internal/provider/client"
	"github.com/stretchr/testify/assert"
)

var (
	version = "2022-06-28"
	token   = "secret_1234567890"
)

func TestNewNotionApiClient(t *testing.T) {

	client := client.NewNotionApiClient()
	assert.Equal(t, "", client.NotionApiVersion)
	assert.Equal(t, "", client.NotionIntegrationToken)

	client.Init(token, version)
	assert.Equal(t, version, client.NotionApiVersion)
	assert.Equal(t, token, client.NotionIntegrationToken)
}

func TestPostSuccessful(t *testing.T) {
	notion_client := client.NewNotionApiClient()
	notion_client.Init(token, version)

	body := `{
		"name": "John Doe",
	}`

	response, err := notion_client.Post("http://0.0.0.0:8000", body)

	if response != nil && response.Body != nil {
		defer response.Body.Close()
		body := make([]byte, 1024)
		n, _ := response.Body.Read(body)
		t.Logf("Response body: %s", body[:n])
	}

	assert.Equal(t, 200, response.StatusCode)
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestPostFails(t *testing.T) {
	notion_client := client.NewNotionApiClient()

	body := `{
		"name": "John Doe",
	}`

	response, err := notion_client.Post("http://invalid-url.com", body)

	assert.Error(t, err)
	assert.Nil(t, response)
}
