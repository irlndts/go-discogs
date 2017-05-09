package main

import (
	"fmt"

	"github.com/irlndts/go-discogs"
)

func main() {
	d := discogs.NewClient().UserAgent("TestDiscogsClient/0.0.1 +http://irlndts.moscow").Token("oQTQKAprakIQfWOkAxTdYyDpgUqahHtdbHTuYkIy")

	if err := d.Currency("EUR"); err != nil {
		fmt.Println(err)
		return
	}

	release, err := d.Release(9893847)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", release)
}
