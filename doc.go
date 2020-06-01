/*
Package discogs is a Go client library for the Discogs API.

The discogs package provides a client for accessing the Discogs API.
First of all import library and init client variable.
According to discogs api documentation you must provide your user-agent.

Some requests require authentification (as any user).
According to Discogs, to send requests with Discogs Auth, you have two options:
sending your credentials in the query string with key and secret parameters or
a token parameter. This is token way example:

	client, err := discogs.New(&discogs.Options{
		UserAgent: "Some Name",
		Currency:  "EUR", // optional, "USD" (default), "GBP", "EUR", "CAD", "AUD", "JPY", "CHF", "MXN", "BRL", "NZD", "SEK", "ZAR" are allowed
		Token:     "Some Token", // optional
		URL:       "https://api.discogs.com", // optional
	})

*/
package discogs
