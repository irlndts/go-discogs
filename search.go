package discogs

import (
	"net/http"

	"github.com/irlndts/go-apirequest"
)

// SearchService ...
type SearchService struct {
	api *apirequest.API
}

// SerachRequest describes search request json
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

// Search describes search response
type Search struct {
	Pagination Page     `json:"pagination"`
	Results    []Result `json:"results,omitempty"`
}

// Result describes a part of search result
type Result struct {
	Style        []string  `json:"style,omitempty"`
	Thumb        string    `json:"thumb,omitempty"`
	Title        string    `json:"title,omitempty"`
	Country      string    `json:"country,omitempty"`
	Format       []string  `json:"format,omitempty"`
	Uri          string    `json:"uri,omitempty"`
	Community    Community `json:"community,omitempty"`
	Label        []string  `json:"label,omitempty"`
	Catno        string    `json:"catno,omitempty"`
	Year         string    `json:"year,omitempty"`
	Genre        []string  `json:"genre,omitempty"`
	Resource_url string    `json:"resource_url,omitempty"`
	Type         string    `json:"type,omitempty"`
	Id           int       `json:"id,omitempty"`
}

func newSearchService(api *apirequest.API) *SearchService {
	return &SearchService{
		api: api.Path("database/search"),
	}
}

// Search makes search request to discogs.
// Issue a search query to our database. This endpoint accepts pagination parameters.
// Authentication (as any user) is required.
// https://www.discogs.com/developers/#page:database,header:database-search
func (self *SearchService) Search(params *SearchRequest) (*Search, *http.Response, error) {
	search := new(Search)
	apiError := new(APIError)

	resp, err := self.api.New().QueryStruct(params).Receive(search, apiError)
	return search, resp, relevantError(err, *apiError)
}
