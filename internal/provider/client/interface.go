package client

import "net/http"

type NotionAPI interface {
	Post(url string, body string) (*http.Response, error)
	Patch(url string, body string) (*http.Response, error)
	Get(url string) (*http.Response, error)
	Init(token string, version string)
}
