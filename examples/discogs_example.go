package main

import (
	"fmt"
	"github.com/irlndts/go-discogs"
	//	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	d := discogs.NewClient(client)
	params := &discogs.ReleaseParams{12345}
	release, resp, err := d.Release.Release(params)
	fmt.Println(release.Title)
	fmt.Println(resp)
	fmt.Println(err)

}
