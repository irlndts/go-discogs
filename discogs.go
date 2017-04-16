package discogs

import (
	"net/http"

	"github.com/irlndts/go-apirequest"
)

const (
	discogsAPI = "https://api.discogs.com/"
)

// Client is a Discogs client for making Discogs API requests.
type Client struct {
	api     *apirequest.API
	Release *ReleaseService
	Master  *MasterService
	Artist  *ArtistService
	Label   *LabelService
	Search  *SearchService
}

// NewClient returns a new Client.
func NewClient(useragent, token string) *Client {
	base := apirequest.New().Client(&http.Client{}).Base(discogsAPI).Add("User-Agent", useragent).Add("Authorization", "Discogs token="+token)
	return &Client{
		Artist:  newArtistService(base.New()),
		Label:   newLabelService(base.New()),
		Master:  newMasterService(base.New()),
		Release: newReleaseService(base.New()),
		Search:  newSearchService(base.New()),
	}
}
