package discogs

import (
	"strconv"
)

// LabelService ...
type LabelService struct {
	url string
}

func newLabelService(url string) *LabelService {
	return &LabelService{
		url: url,
	}
}

// Label resource represents a label, company, recording studio, location,
// or other entity involved with artists and releases.
type Label struct {
	Profile     string     `json:"profile"`
	ReleasesURL string     `json:"releases_url"`
	Name        string     `json:"name"`
	ContactInfo string     `json:"contact_info"`
	URI         string     `json:"uri"`
	Sublabels   []Sublable `json:"sublabels"`
	URLs        []string   `json:"urls"`
	Images      []Image    `json:"images"`
	ResourceURL string     `json:"resource_url"`
	ID          int        `json:"id"`
	DataQuality string     `json:"data_quality"`
}

// Label returns a label.
func (s *LabelService) Label(labelID int) (*Label, error) {
	var label *Label
	if err := request(s.url+strconv.Itoa(labelID), nil, &label); err != nil {
		return nil, err
	}
	return label, nil
}

// LabelReleases is a list of Releases associated with the label.
type LabelReleases struct {
	Pagination Page            `json:"pagination"`
	Releases   []ReleaseSource `json:"releases"`
}

// Releases returns a list of Releases associated with the label.
func (s *LabelService) Releases(labelID int, pagination *Pagination) (*LabelReleases, error) {
	var releases *LabelReleases
	if err := request(s.url+strconv.Itoa(labelID)+"/releases", pagination.toParams(), &releases); err != nil {
		return nil, err
	}
	return releases, nil
}
