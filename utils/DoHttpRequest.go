package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func DoHttpRequest(parsedUrl *url.URL, allowedHeaders []string) (int, map[string]string, []byte, error) {
	// Do a http request to that URL
	resp, err := http.Get(parsedUrl.String())
	if err != nil {
		return 0, nil, nil, err
	}
	// Read the response Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, nil, err
	}
	// Extract the allowed Headers
	headers := make(map[string]string, 0)
	for header, value := range resp.Header {
		for _, allowHeader := range allowedHeaders {
			if strings.ToLower(allowHeader) == strings.ToLower(header) {
				headers[header] = value[0]
				break
			}
		}
	}
	return resp.StatusCode, headers, body, nil
}
