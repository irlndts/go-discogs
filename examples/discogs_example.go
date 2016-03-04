package main

import (
	"fmt"
	"github.com/irlndts/go-discogs/discogs"
	"net/http"
)

func main() {
	client := &http.Client{}
	discogs := discogs.NewClient(client)
	fmt.Println(discogs)

}
