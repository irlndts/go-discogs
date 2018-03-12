package discogs

import (
	"net/url"
	"strconv"
)

// Release serves relesase response from discogs
type Release struct {
	Title             string         `json:"title"`
	ID                int            `json:"id"`
	Artists           []ArtistSource `json:"artists"`
	DataQuality       string         `json:"data_quality"`
	Thumb             string         `json:"thumb"`
	Community         Community      `json:"community"`
	Companies         []Company      `json:"companies"`
	Country           string         `json:"country"`
	DateAdded         string         `json:"date_added"`
	DateChanged       string         `json:"date_changed"`
	EstimatedWeight   int            `json:"estimated_weight"`
	ExtraArtists      []ArtistSource `json:"extraartists"`
	FormatQuantity    int            `json:"format_quantity"`
	Formats           []Format       `json:"formats"`
	Genres            []string       `json:"genres"`
	Identifiers       []Identifier   `json:"identifiers"`
	Images            []Image        `json:"images"`
	Labels            []LabelSource  `json:"labels"`
	LowestPrice       float64        `json:"lowest_price"`
	MasterID          int            `json:"master_id"`
	MasterURL         string         `json:"master_url"`
	Notes             string         `json:"notes,omitempty"`
	NumForSale        int            `json:"num_for_sale,omitempty"`
	Released          string         `json:"released"`
	ReleasedFormatted string         `json:"released_formatted"`
	ResourceURL       string         `json:"resource_url"`
	// Series
	Status    string   `json:"status"`
	Styles    []string `json:"styles"`
	Tracklist []Track  `json:"tracklist"`
	URI       string   `json:"uri"`
	Videos    []Video  `json:"videos"`
	Year      int      `json:"year"`
}

type ReqRelease struct {
	CurrAbbr string
}

// ReleaseService ...
type ReleaseService struct {
	url      string
	currency string
}

func newReleaseService(url string, currency string) *ReleaseService {
	return &ReleaseService{
		url:      url,
		currency: currency,
	}
}

// Release returns release by release's ID
func (s *ReleaseService) Release(releaseID int) (*Release, error) {
	params := url.Values{}
	params.Set("CurrAbbr", s.currency)

	var release *Release
	if err := request(s.url+strconv.Itoa(releaseID), params, &release); err != nil {
		return nil, err
	}

	return release, nil
}
