package main

import (
	"fmt"

	"github.com/irlndts/go-discogs"
)

func main() {
	d, err := discogs.NewClient(&discogs.Options{
		UserAgent: "TestDiscogsClient/0.0.1 +http://example.com",
		Currency:  "USD",
		Token:     "",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	master, err := d.Database.Master(718441)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", master)
}
