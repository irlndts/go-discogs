package discogs

import (
	"github.com/irlndts/go-apirequest"
	"net/http"
)

const (
	discogsAPI = "https://api.discogs.com/"
	useragent  = "Test UserAgent"
)

// Client is a Discogs client for making Discogs API requests.
type Client struct {
	api     *apirequest.API
	Release *ReleaseService
	Master  *MasterService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := apirequest.New().Client(httpClient).Base(discogsAPI).Add("User-Agent", useragent)

	return &Client{
		api:     base,
		Release: newReleaseService(base.New()),
		Master:  newMasterService(base.New()),
	}
}

// discogs require specified user agent
func (c *Client) UserAgent(useragent string) *Client {
	c.api.Set("User-Agent", useragent)
	return c
}
