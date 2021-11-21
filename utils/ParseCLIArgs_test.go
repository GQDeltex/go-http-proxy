package utils

import (
	"testing"
	"time"
)

func TestDefaultValues(t *testing.T) {
	cliArgs, err := ParseCLIArgs()
	if err != nil {
		t.Error(err.Error())
	}
	if cliArgs.LogLevel.String() != "warning" {
		t.Errorf(`Loglevel: wanted: "warning" got: "%s"`, cliArgs.LogLevel.String())
	}
	reference_allowedHeaders := make([]string, 0)
	reference_allowedHeaders = append(reference_allowedHeaders, "Content-Type")
	if len(cliArgs.AllowedHeaders) != len(reference_allowedHeaders) {
		t.Errorf("AllowedHeaders: wanted: len %d got: len %d", len(reference_allowedHeaders), len(cliArgs.AllowedHeaders))
	} else {
		for _, header := range cliArgs.AllowedHeaders {
			found := false
			for _, ref_header := range reference_allowedHeaders {
				if header == ref_header {
					found = true
				}
			}
			if found == false {
				t.Errorf(`AllowedHeaders: wanted: "%s" got: ""`, header)
				break
			}
		}
	}
	if cliArgs.Port != 3000 {
		t.Errorf("Port: wanted: 3000 got: %d", cliArgs.Port)
	}
	if cliArgs.CachingTime != time.Minute*30 {
		t.Errorf(`CachingTime: wanted: "30m0s" got: "%s"`, cliArgs.CachingTime.String())
	}
	if cliArgs.Secret != "secret" {
		t.Errorf(`Secret: wanted: "secret" got: "%s"`, cliArgs.Secret)
	}
	if cliArgs.UserAgent != "go-http-proxy" {
		t.Errorf(`UserAgent: wanted: "go-http-proxy" got: "%s"`, cliArgs.UserAgent)
	}
}
