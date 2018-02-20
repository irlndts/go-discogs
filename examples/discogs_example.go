package main

import (
	"fmt"

	"github.com/irlndts/go-discogs"
)

func main() {
	d, err := discogs.NewClient(&discogs.Options{
		UserAgent: "TestDiscogsClient/0.0.1 +http://example.com",
		Currency:  "EUR",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	release, err := d.Release.Release(9893847)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", release)
}
