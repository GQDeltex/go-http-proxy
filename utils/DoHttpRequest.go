package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func DoHttpRequest(parsedUrl *url.URL) (int, string, []byte, error) {
	// Do a http request to that URL
	resp, err := http.Get(parsedUrl.String())
	if err != nil {
		return 0, "", nil, err
	}
	// Read the response Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", nil, err
	}
	// Extract the content-type
	content_type := resp.Header["Content-Type"][0]
	content_type = strings.Split(content_type, "/")[1]
	return resp.StatusCode, content_type, body, nil
}
