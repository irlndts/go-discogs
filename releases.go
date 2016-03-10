package discogs

import (
	"github.com/irlndts/go-apirequest"
	"net/http"
)

type ReleaseService struct {
	api *apirequest.API
}

type ReleaseParams struct {
	Release_id string
}

type Release struct {
	Title              string         `json:"title"`
	Id                 int            `json:"id"`
	Artists            []ArtistSource `json:"artists"`
	Data_quality       string         `json:"data_quality"`
	Thumb              string         `json:"thumb"`
	Community          Community      `json:"community"`
	Companies          []Company      `json:"companies"`
	Country            string         `json:"country"`
	Date_added         string         `json:"date_added"`
	Date_changed       string         `json:"date_changed"`
	Estimated_weight   int            `json:"estimated_weight"`
	Extraartists       []ArtistSource `json:"extraartists"`
	Format_quantity    int            `json:"format_quantity"`
	Formats            []Format       `json:"formats"`
	Genres             []string       `json:"genres"`
	Identifiers        []Identifier   `json:"identifiers"`
	Images             []Image        `json:"images"`
	Labels             []LabelSource  `json:"labels"`
	Master_id          int            `json:"master_id"`
	Master_url         string         `json:"master_url"`
	Notes              string         `josn:"notes"`
	Released           string         `json:"released"`
	Released_formatted string         `json:"released_formatted"`
	Resource_url       string         `json:"resource_url"`
	Status             string         `json:"status"`
	Styles             []string       `json:"styles"`
	Tracklist          []Track        `json:"tracklist"`
	Uri                string         `json:"uri"`
	Videos             []Video        `json:"videos"`
	Year               int            `json:"year"`
}

func newReleaseService(api *apirequest.API) *ReleaseService {
	return &ReleaseService{
		api: api.Path("releases/"),
	}
}

func (self *ReleaseService) Release(params *ReleaseParams) (*Release, *http.Response, error) {
	release := new(Release)
	apiError := new(APIError)

	resp, err := self.api.New().Get(params.Release_id).Receive(release, apiError)
	return release, resp, relevantError(err, *apiError)
}
