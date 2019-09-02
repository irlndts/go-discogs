package discogs

import (
	"net/url"
	"strconv"
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

// SearchRequest describes search request
type SearchRequest struct {
	Q            string // search query
	Type         string // one of release, master, artist, label
	Title        string // search by combined “Artist Name - Release Title” title field
	ReleaseTitle string // search release titles
	Credit       string // search release credits
	Artist       string // search artist names
	Anv          string // search artist ANV
	Label        string // search label names
	Genre        string // search genres
	Style        string // search styles
	Country      string // search release country
	Year         string // search release year
	Format       string // search formats
	Catno        string // search catalog number
	Barcode      string // search barcodes
	Track        string // search track titles
	Submitter    string // search submitter username
	Contributor  string // search contributor usernames

	Page    int
	PerPage int
}

func (r *SearchRequest) params() url.Values {
	if r == nil {
		return nil
	}

	params := url.Values{}
	params.Set("q", r.Q)
	params.Set("type", r.Type)
	params.Set("title", r.Title)
	params.Set("release_title", r.ReleaseTitle)
	params.Set("credit", r.Credit)
	params.Set("artist", r.Artist)
	params.Set("anv", r.Anv)
	params.Set("label", r.Label)
	params.Set("genre", r.Genre)
	params.Set("style", r.Style)
	params.Set("country", r.Country)
	params.Set("year", r.Year)
	params.Set("format", r.Format)
	params.Set("catno", r.Catno)
	params.Set("barcode", r.Barcode)
	params.Set("track", r.Track)
	params.Set("submitter", r.Submitter)
	params.Set("contributor", r.Contributor)
	params.Set("page", strconv.Itoa(r.Page))
	params.Set("per_page", strconv.Itoa(r.PerPage))
	return params
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
	CoverImage  string    `json:"cover_image,omitempty"`
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
func (s *SearchService) Search(req SearchRequest) (*Search, error) {
	var search *Search
	err := request(s.url, req.params(), &search)
	return search, err
}
