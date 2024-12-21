package client

import (
	"net/http"
	"sync"
)

type NotionApiClient struct {
	Client                 *http.Client
	NotionApiVersion       string
	NotionIntegrationToken string
}

var _ NotionAPI = (*NotionApiClient)(nil)

var (
	instance *NotionApiClient
	once     sync.Once
)

func NewNotionApiClient() *NotionApiClient {
	once.Do(func() {
		instance = &NotionApiClient{
			Client:                 &http.Client{},
			NotionApiVersion:       "",
			NotionIntegrationToken: "",
		}
	})
	return instance
}

func (c *NotionApiClient) Init(token string, version string) {
	c.NotionIntegrationToken = token
	c.NotionApiVersion = version
}
