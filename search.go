package discogs

import (
	"github.com/google/go-querystring/query"
)

// SearchService ...
type SearchService struct {
	url string
}

func newSearchService(url string) *SearchService {
	return &SearchService{
		url: url,
	}
}

// SerachRequest describes search request
type SearchRequest struct {
	Q            string `url:"q,omitempty"`             // search query
	Type         string `url:"type,omitempty"`          // one of release, master, artist, label
	Title        string `url:"title,omitempty"`         // search by combined “Artist Name - Release Title” title field
	ReleaseTitle string `url:"release_title,omitempty"` // search release titles
	Credit       string `url:"credit,omitempty"`        // search release credits
	Artist       string `url:"artist,omitempty"`        // search artist names
	Anv          string `url:"anv,omitempty"`           // search artist ANV
	Label        string `url:"label,omitempty"`         // search label names
	Genre        string `url:"genre,omitempty"`         // search genres
	Style        string `url:"style,omitempty"`         // search styles
	Country      string `url:"country,omitempty"`       // search release country
	Year         string `url:"year,omitempty"`          // search release year
	Format       string `url:"format,omitempty"`        // search formats
	Catno        string `url:"catno,omitempty"`         // search catalog number
	Barcode      string `url:"barcode,omitempty"`       // search barcodes
	Track        string `url:"track,omitempty"`         // search track titles
	Submitter    string `url:"submitter,omitempty"`     // search submitter username
	Contributer  string `url:"contributer,omitempty"`   // search contributor usernames

	Page    int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
}

// Search describes search response
type Search struct {
	Pagination Page     `json:"pagination"`
	Results    []Result `json:"results,omitempty"`
}

// Result describes a part of search result
type Result struct {
	Style       []string  `json:"style,omitempty"`
	Thumb       string    `json:"thumb,omitempty"`
	Title       string    `json:"title,omitempty"`
	Country     string    `json:"country,omitempty"`
	Format      []string  `json:"format,omitempty"`
	URI         string    `json:"uri,omitempty"`
	Community   Community `json:"community,omitempty"`
	Label       []string  `json:"label,omitempty"`
	Catno       string    `json:"catno,omitempty"`
	Year        string    `json:"year,omitempty"`
	Genre       []string  `json:"genre,omitempty"`
	ResourceURL string    `json:"resource_url,omitempty"`
	Type        string    `json:"type,omitempty"`
	ID          int       `json:"id,omitempty"`
}

// Search makes search request to discogs.
// Issue a search query to our database. This endpoint accepts pagination parameters.
// Authentication (as any user) is required.
// https://www.discogs.com/developers/#page:database,header:database-search
// TODO(irlndts): improve params to pass
func (s *SearchService) Search(req SearchRequest) (*Search, error) {
	params, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	var search *Search
	if err := request(s.url, params, &search); err != nil {
		return nil, err
	}
	return search, nil
}
