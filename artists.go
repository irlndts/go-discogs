package discogs

import (
	"net/http"

	"github.com/irlndts/go-apirequest"
)

// ArtistService ...
type ArtistService struct {
	api *apirequest.API
}

// ArtistParams ...
type ArtistParams struct {
	Artist_id  string
	Sort       string // year, title, format
	Sort_order string // asc, desc
	Page       int
	Per_page   int
}

// Artist ...
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

<<<<<<< HEAD
// ArtistReleases ...
=======
// Artistreleases ...
>>>>>>> efff71f46e36d0fd1868f4abe3efd8c82e32386f
type ArtistReleases struct {
	Paginastion Page            `json:"pagination"`
	Releases    []ReleaseSource `json:"releases"`
}

func newArtistService(api *apirequest.API) *ArtistService {
	return &ArtistService{
		api: api.Path("artists/"),
	}
}

// Artist ...
func (self *ArtistService) Artist(params *ArtistParams) (*Artist, *http.Response, error) {
	artist := new(Artist)
	apiError := new(APIError)

	resp, err := self.api.New().Get(params.Artist_id).Receive(artist, apiError)
	return artist, resp, relevantError(err, *apiError)
}

// Releases ...
func (self *ArtistService) Releases(params *ArtistParams) (*ArtistReleases, *http.Response, error) {
	releases := new(ArtistReleases)
	apiError := new(APIError)

	resp, err := self.api.New().Get(params.Artist_id+"/releases").QueryStruct(params).Receive(releases, apiError)
	return releases, resp, relevantError(err, *apiError)
}
