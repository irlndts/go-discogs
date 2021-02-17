package discogs

import (
	"testing"
)

const (
	testUserAgent = "UnitTestClient/0.0.2"
	testUsername  = "test_user"
	testToken     = ""
)

func initDiscogsClient(t *testing.T, options *Options) Discogs {
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

	client, err := New(options)
	if err != nil {
		t.Fatalf("failed to create client: %s", err)
	}

	return client
}

func TestNew(t *testing.T) {
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

	for name := range tests {
		tt := tests[name]
		t.Run(name, func(t *testing.T) {
			if _, err := New(tt.options); err != tt.err {
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
		{currency: "", want: "USD"},
		{currency: "USD", want: "USD"},
		{currency: "GBP", want: "GBP"},
		{currency: "EUR", want: "EUR"},
		{currency: "CAD", want: "CAD"},
		{currency: "AUD", want: "AUD"},
		{currency: "JPY", want: "JPY"},
		{currency: "CHF", want: "CHF"},
		{currency: "MXN", want: "MXN"},
		{currency: "BRL", want: "BRL"},
		{currency: "NZD", want: "NZD"},
		{currency: "SEK", want: "SEK"},
		{currency: "ZAR", want: "ZAR"},
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
