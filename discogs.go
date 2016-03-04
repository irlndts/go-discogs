package discogs

import (
	"github.com/irlndts/go-apirequest"
	"net/http"
)

const (
	discogsAPI = "https://api.discogs.com/"
	useragent  = "TestDiscogsClient/0.0.1 +http://irlndts.moscow"
)

// Client is a Discogs client for making Discogs API requests.
type Client struct {
	api     *apirequest.API
	Release *ReleaseService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := apirequest.New().Client(httpClient).Base(discogsAPI).Add("User-Agent", useragent)

	return &Client{
		api:     base,
		Release: newReleaseService(base.New()),
	}
}
