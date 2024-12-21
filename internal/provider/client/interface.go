package client

import "net/http"

type NotionAPI interface {
	Init(token string, version string)

	Post(url string, body string) (*http.Response, error)
	Patch(url string, body string) (*http.Response, error)
	Get(url string) (*http.Response, error)

	HandleApiResponse(response *http.Response, err error, baseMessage string, addError func(string, string)) bool
}
