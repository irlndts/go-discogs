package main

import (
	"fmt"
	"github.com/irlndts/go-discogs"
	"net/http"
)

func main() {
	client := &http.Client{}
	d := discogs.NewClient(client).UserAgent("TestDiscogsClient/0.0.1 +http://irlndts.moscow")

	/*
		params := &discogs.ReleaseParams{Release_id: "8138518"}
		release, _, err := d.Release.Release(params)
	*/
	params := &discogs.LabelParams{Label_id: "1000"}
	label, _, err := d.Label.Label(params)

	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(release.Title)
		fmt.Println(label.Name)
	}
	//fmt.Println(resp)
}
