package utils

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
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
	// Remove relative link requests
	if parsedUrl.Hostname() == "" {
		return nil, errors.New("No hostname was given")
	}
	// Allow only http:// or https:// requests
	if parsedUrl.Scheme != "https" && parsedUrl.Scheme != "http" {
		return nil, errors.New("Unsupported protocol")
	}
	// block localhost and 127.0.0.1
	if parsedUrl.Hostname() == "localhost" {
		return nil, errors.New("Disallow Private Subnets")
	}
	if parsedUrl.Hostname() == "127.0.0.1" {
		return nil, errors.New("Disallow Private Subnets")
	}
	// block 10.0.0.0/8
	if strings.HasPrefix(parsedUrl.Hostname(), "10.") {
		return nil, errors.New("Disallow Private Subnets")
	}
	// block 172.16.0.0/12
	for i := 16; i <= 32; i++ {
		if strings.HasPrefix(parsedUrl.Hostname(), "172."+strconv.Itoa(i)+".") {
			return nil, errors.New("Disallow Private Subnets")
		}
	}
	// block 192.168.0.0/16
	if strings.HasPrefix(parsedUrl.Hostname(), "192.168.") {
		return nil, errors.New("Disallow Private Subnets")
	}
	return parsedUrl, nil
}
