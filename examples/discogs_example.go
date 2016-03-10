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
	params := &discogs.MasterVersionParams{Master_id: "1000", Page: 2, Per_page: 1}
	master, _, err := d.Master.Versions(params)

	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(release.Title)
		fmt.Println(master)
	}
	//fmt.Println(resp)
}
