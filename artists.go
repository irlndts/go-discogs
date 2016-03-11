package discogs

import (
	"github.com/irlndts/go-apirequest"
	"net/http"
)

type ArtistService struct {
	api *apirequest.API
}

type ArtistParams struct {
	Artist_id  string
	Sort       string // year, title, format
	Sort_order string // asc, desc
	Page       int
	Per_page   int
}

type Artist struct {
	Namevariations []string `json:"namevariations"`
	Profile        string   `json:"profile"`
	Releases_url   string   `json:"releases_url"`
	Resource_url   string   `json:"resource_url"`
	Uri            string   `json:"uri"`
	Urls           []string `json:"urls"`
	Data_quality   string   `json:"data_quality"`
	Id             int      `json:"id"`
	Images         []Image  `json:"images"`
	Members        []Member `json:"members"`
}

type ArtistReleases struct {
	Paginastion Page            `json:"pagination"`
	Releases    []ReleaseSource `json:"releases"`
}

func newArtistService(api *apirequest.API) *ArtistService {
	return &ArtistService{
		api: api.Path("artists/"),
	}
}

func (self *ArtistService) Artist(params *ArtistParams) (*Artist, *http.Response, error) {
	artist := new(Artist)
	apiError := new(APIError)

	resp, err := self.api.New().Get(params.Artist_id).Receive(artist, apiError)
	return artist, resp, relevantError(err, *apiError)
}

func (self *ArtistService) Releases(params *ArtistParams) (*ArtistReleases, *http.Response, error) {
	releases := new(ArtistReleases)
	apiError := new(APIError)

	resp, err := self.api.New().Get(params.Artist_id+"/releases").QueryStruct(params).Receive(releases, apiError)
	return releases, resp, relevantError(err, *apiError)
}
