package utils

import (
	"errors"
	"net/url"
)

func ParseURL(urlstr string) (*url.URL, error) {
	if urlstr == "" {
		return nil, errors.New("No URl parameter")
	}
	// Check the validity of the url
	parsedUrl, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	if parsedUrl.Hostname() == "" {
		return nil, errors.New("No hostname was given")
	}
	return parsedUrl, nil
}
