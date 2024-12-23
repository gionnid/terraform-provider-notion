package client

import (
	"bytes"
	"net/http"
	"time"
)

func (c *NotionApiClient) GenericRequest(url string, method string, body string) (*http.Response, error) {
	waiter := NewWaiter(c, 340*time.Millisecond)
	waiter.WaitToReserveSpot()

	headers := c.GetHeaders(true)
	if body == "" {
		headers = c.GetHeaders(false)
	}

	bufferBody := bytes.NewBuffer([]byte(body))

	req, err := http.NewRequest(method, url, bufferBody)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := c.Client.Do(req)

	waiter.ReleaseSpot()
	return resp, err
}

func (c *NotionApiClient) Post(url string, body string) (*http.Response, error) {
	return c.GenericRequest(url, "POST", body)
}

func (c *NotionApiClient) Patch(url string, body string) (*http.Response, error) {
	return c.GenericRequest(url, "PATCH", body)
}

func (c *NotionApiClient) Get(url string) (*http.Response, error) {
	return c.GenericRequest(url, "GET", "")
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
