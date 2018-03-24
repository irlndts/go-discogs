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

func TestNewClient(t *testing.T) {
	tests := map[string]struct {
		options *Options
		err     error
	}{
		"normal": {&Options{
			UserAgent: testUserAgent,
			Currency:  "USD",
			Token:     "some token",
		}, nil},
		"incorrect user-agent": {&Options{
			UserAgent: "",
			Currency:  "USD",
		}, ErrUserAgentInvalid},
		"incorrect currency": {&Options{
			UserAgent: testUserAgent,
			Currency:  "RUR",
		}, ErrCurrencyNotSupported},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := NewClient(tt.options); err != tt.err {
				t.Errorf("err got=%s; want=%s", err, tt.err)
			}
		})
	}
}

func TestCurrency(t *testing.T) {
	tests := []struct {
		currency string
		want     string
		err      error
	}{
		{currency: "", want: "USD", err: nil},
		{currency: "USD", want: "USD", err: nil},
		{currency: "GBP", want: "GBP", err: nil},
		{currency: "EUR", want: "EUR", err: nil},
		{currency: "CAD", want: "CAD", err: nil},
		{currency: "AUD", want: "AUD", err: nil},
		{currency: "JPY", want: "JPY", err: nil},
		{currency: "CHF", want: "CHF", err: nil},
		{currency: "MXN", want: "MXN", err: nil},
		{currency: "BRL", want: "BRL", err: nil},
		{currency: "NZD", want: "NZD", err: nil},
		{currency: "SEK", want: "SEK", err: nil},
		{currency: "ZAR", want: "ZAR", err: nil},
		{currency: "RUR", want: "", err: ErrCurrencyNotSupported},
	}
	for i, tt := range tests {
		cur, err := currency(tt.currency)
		if err != tt.err {
			t.Errorf("#%d err got=%s; want=%s", i, err, tt.err)
		}
		if cur != tt.want {
			t.Errorf("#%d currency got=%s; want=%s", i, cur, tt.want)
		}
	}
}
