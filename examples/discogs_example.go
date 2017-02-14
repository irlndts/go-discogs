package main

import (
	"fmt"

	"github.com/irlndts/go-discogs"
)

func main() {
	d := discogs.NewClient("TestDiscogsClient/0.0.1 +http://irlndts.moscow", "")

	request := &discogs.SearchRequest{Q: "The Reggaenauts - River Rock / Thursday Kick-off", Page: 0, Per_page: 1}
	search, _, err := d.Search.Search(request)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range search.Results {
		fmt.Println(r.Id, r.Title)
	}
}
