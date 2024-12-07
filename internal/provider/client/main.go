package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NotionApiClient represents a client for making HTTP requests
// Implemented as a singleton
type NotionApiClient struct {
	Client                 *http.Client
	NotionApiVersion       string
	NotionIntegrationToken string
}

var (
	instance *NotionApiClient
	once     sync.Once
)

func GetNotionApiClient() *NotionApiClient {
	return instance
}

func NewNotionApiClient(ctx context.Context, version string, token string) *NotionApiClient {
	once.Do(func() {
		instance = &NotionApiClient{
			Client:                 &http.Client{},
			NotionApiVersion:       version,
			NotionIntegrationToken: token,
		}
	})

	tflog.Info(ctx, "Notion API client initialized.")

	return instance
}

func (c *NotionApiClient) GetHeaders(include_content_type bool) map[string]string {
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + c.NotionIntegrationToken
	headers["Notion-Version"] = c.NotionApiVersion
	if include_content_type {
		headers["Content-Type"] = "application/json"
	}
	return headers
}

// PostRequest performs a generic POST request
func (c *NotionApiClient) Post(url string, headers map[string]string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.Client.Do(req)
}
