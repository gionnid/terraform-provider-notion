package utility

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// NotionApiClient represents a client for making HTTP requests
type NotionApiClient struct {
	Client                 *http.Client
	NotionApiVersion       string
	NotionIntegrationToken string
}

// NewNotionApiClient creates a new HTTPClient with a default http.Client
func NewNotionApiClient(version string, token string) *NotionApiClient {
	return &NotionApiClient{
		Client:                 &http.Client{},
		NotionApiVersion:       version,
		NotionIntegrationToken: token,
	}
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
