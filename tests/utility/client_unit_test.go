package utility_test

import (
	"testing"

	"github.com/gionnid/terraform-provider-notion/internal/provider/utility"
	"github.com/stretchr/testify/assert"
)

func TestNewNotionApiClient(t *testing.T) {
	version := "2022-06-28"
	token := "secret_1234567890"
	client := utility.NewNotionApiClient(version, token)
	assert.NotNil(t, client)
	assert.Equal(t, version, client.NotionApiVersion)
}

func TestPost(t *testing.T) {
	client := utility.NewNotionApiClient("2022-06-28", "secret_1234567890")

	headers := client.GetHeaders(true)

	body := map[string]interface{}{
		"name": "John Doe",
	}

	response, err := client.Post("http://0.0.0.0:8000", headers, body)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}
