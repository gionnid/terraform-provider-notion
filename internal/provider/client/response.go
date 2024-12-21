package client

import (
	"io"
	"net/http"
	"strconv"
)

func (c *NotionApiClient) HandleApiResponse(response *http.Response, err error, baseMessage string, addError func(string, string)) (can_continue bool) {
	if err != nil {
		addError(baseMessage, err.Error())
		return false
	}
	if response.StatusCode != 200 {
		body, _ := io.ReadAll(response.Body)
		addError(baseMessage, "Status Code: "+strconv.Itoa(response.StatusCode)+" Message: "+string(body))
		return false
	}
	return true
}
