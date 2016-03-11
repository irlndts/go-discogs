package main

import (
	"fmt"
	"github.com/irlndts/go-discogs"
)

func main() {
	d := discogs.NewClient().UserAgent("TestDiscogsClient/0.0.1 +http://irlndts.moscow")

	/*
		params := &discogs.ReleaseParams{Release_id: "8138518"}
		release, _, err := d.Release.Release(params)
	*/
	params := &discogs.LabelParams{Label_id: "890477", Page: 2, Per_page: 3}
	label, _, err := d.Label.Releases(params)

	if err != nil {
		fmt.Println(err)
	} else {
		for _, release := range label.Releases {
			fmt.Println(release.Title)
		}
	}
}
