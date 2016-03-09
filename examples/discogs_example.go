package main

import (
	"fmt"
	"github.com/irlndts/go-discogs"
	//	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	d := discogs.NewClient(client).UserAgent("TestDiscogsClient/0.0.1 +http://irlndts.moscow")
	params := &discogs.ReleaseParams{Release_id: "1"}
	release, _, err := d.Release.Release(params)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(release)
	}
	//fmt.Println(resp)
}
