package discogs

import (
	"fmt"
	"github.com/irlndts/go-apirequest"
	"net/http"
)

const (
	discogsAPI = "https://api.discogs.com/"
	useragent  = "TestDiscogsClient/0.0.1 +http://irlndts.moscow"
)

// Client is a Discogs client for making Discogs API requests.
type Client struct {
	api      *apirequest.API
	Releases *ReleaseService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := apirequest.New().Client(httpClient).Base(discogsAPI).Add("User-Agent", useragent)

	fmt.Println(base)

	return &Client{
		api:      base,
		Releases: newReleaseService(base.New()),
	}
}
