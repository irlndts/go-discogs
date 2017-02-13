package main

import (
	"fmt"

	"github.com/irlndts/go-discogs"
)

func main() {
	d := discogs.NewClient().UserAgent("TestDiscogsClient/0.0.1 +http://irlndts.moscow").Token("")

	params := &discogs.SearchRequest{Release_title: "nevermind", Artist: "nirvana"}
	search, _, err := d.Search.Search(params)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range search.Results {
		fmt.Println(r)
	}
}
