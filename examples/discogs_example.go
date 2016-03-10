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
	params := &discogs.ArtistParams{Artist_id: "1000"}
	artist, _, err := d.Artist.Artist(params)

	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(release.Title)
		fmt.Println(artist.Id)
	}
	//fmt.Println(resp)
}
