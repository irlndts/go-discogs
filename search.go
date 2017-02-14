package discogs

import (
	"net/http"

	"github.com/irlndts/go-apirequest"
)

type SearchService struct {
	api *apirequest.API
}

type SearchRequest struct {
	Q             string // search query (optional)
	Type          string // one of release, master, artist, label (optional)
	Title         string // search by combined “Artist Name - Release Title” title field (optional)
	Release_title string // search release titles (optional)
	Credit        string // search release credits (optional)
	Artist        string // search artist names (optional)
	Anv           string // search artist ANV (optional)
	Label         string // search label names (optional)
	Genre         string // search genres (optional)
	Style         string // search styles (optional)
	Country       string // search release country (optional)
	Year          string // search release year (optional)
	Format        string // search formats (optional)
	Catno         string // search catalog number (optional)
	Barcode       string // search barcodes (optional)
	Track         string // search track titles (optional)
	Submitter     string // search submitter username (optional)
	Contributer   string // search contributor usernames (optional)

	Page     int // optional
	Per_page int // optional
}

type Search struct {
	Pagination Page     `json:"pagination"`
	Results    []Result `json:"results"`
}

type Result struct {
	Style        []string  `json:"style"`
	Thumb        string    `json:"thumb"`
	Title        string    `json:"title"`
	Country      string    `json:"country"`
	Format       []string  `json:"format"`
	Uri          string    `json:"uri"`
	Community    Community `json:"community"`
	Label        []string  `json:"label"`
	Catno        string    `json:"catno"`
	Year         string    `json:"year"`
	Genre        []string  `json:"genre"`
	Resource_url string    `json:"resource_url"`
	Type         string    `json:"type"`
	Id           int       `json:"id"`
}

func newSearchService(api *apirequest.API) *SearchService {
	return &SearchService{
		api: api.Path("database/search"),
	}
}

func (self *SearchService) Search(params *SearchRequest) (*Search, *http.Response, error) {
	search := new(Search)
	apiError := new(APIError)

	resp, err := self.api.New().QueryStruct(params).Receive(search, apiError)
	return search, resp, relevantError(err, *apiError)
}
