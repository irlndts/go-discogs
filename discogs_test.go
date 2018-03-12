package discogs

import (
	"testing"
)

const (
	testUserAgent = "UnitTestClient/0.0.2 +https://github.com/irlndts/go-discogs"
	testToken     = ""
)

func initDiscogsClient(t *testing.T, options *Options) *Client {
	if options == nil {
		options = &Options{
			UserAgent: testUserAgent,
			Currency:  "USD",
			Token:     testToken,
		}
	}

	if options.UserAgent == "" {
		options.UserAgent = testUserAgent
	}

	client, err := NewClient(options)
	if err != nil {
		t.Fatalf("failed to create client: %s", err)
	}

	return client
}
