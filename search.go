package discogs

import (
	"net/url"
	"strconv"
)

// SearchService is an interface to work with search.
type SearchService interface {
	// Search makes search request to discogs.
	// Issue a search query to database. This endpoint accepts pagination parameters.
	// Authentication (as any user) is required.
	// https://www.discogs.com/developers/#page:database,header:database-search
	Search(req SearchRequest) (*Search, error)
}

// searchService ...
type searchService struct {
	url string
}

func newSearchService(url string) SearchService {
	return &searchService{
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

	if r.Q != "" {
		params.Set("q", r.Q)
	}
	if r.Type != "" {
		params.Set("type", r.Type)
	}
	if r.Title != "" {
		params.Set("title", r.Title)
	}
	if r.ReleaseTitle != "" {
		params.Set("release_title", r.ReleaseTitle)
	}
	if r.Credit != "" {
		params.Set("credit", r.Credit)
	}
	if r.Artist != "" {
		params.Set("artist", r.Artist)
	}
	if r.Anv != "" {
		params.Set("anv", r.Anv)
	}
	if r.Label != "" {
		params.Set("label", r.Label)
	}
	if r.Genre != "" {
		params.Set("genre", r.Genre)
	}
	if r.Style != "" {
		params.Set("style", r.Style)
	}
	if r.Country != "" {
		params.Set("country", r.Country)
	}
	if r.Year != "" {
		params.Set("year", r.Year)
	}
	if r.Format != "" {
		params.Set("format", r.Format)
	}
	if r.Catno != "" {
		params.Set("catno", r.Catno)
	}
	if r.Barcode != "" {
		params.Set("barcode", r.Barcode)
	}
	if r.Track != "" {
		params.Set("track", r.Track)
	}
	if r.Submitter != "" {
		params.Set("submitter", r.Submitter)
	}
	if r.Contributor != "" {
		params.Set("contributor", r.Contributor)
	}

	params.Set("page", strconv.Itoa(r.Page))
	if r.PerPage != 0 {
		params.Set("per_page", strconv.Itoa(r.PerPage))
	}
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
	MasterID    int       `json:"master_id,omitempty"`
}

func (s *searchService) Search(req SearchRequest) (*Search, error) {
	var search *Search
	err := request(s.url, req.params(), &search)
	return search, err
}
