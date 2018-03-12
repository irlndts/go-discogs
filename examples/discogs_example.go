package main

import (
	"fmt"
	"net/url"

	"github.com/irlndts/go-discogs"
)

func main() {
	d, err := discogs.NewClient(&discogs.Options{
		UserAgent: "TestDiscogsClient/0.0.1 +http://example.com",
		Currency:  "EUR",
		Token:     "",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	params := url.Values{}
	params.Set("q", "Ska-Jazz Review")
	release, err := d.Search.Search(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", release)
}
