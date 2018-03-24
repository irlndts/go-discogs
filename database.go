package discogs

import (
	"net/url"
	"strconv"
)

const (
	releasesURI = "/releases/"
	artistsURI  = "/artists/"
	labelsURI   = "/labels/"
	mastersURI  = "/masters/"
)

// DatabaseService ...
type DatabaseService struct {
	url      string
	currency string
}

func newDatabaseService(url string, currency string) *DatabaseService {
	return &DatabaseService{
		url:      url,
		currency: currency,
	}
}

// ReqRelease serves release request
type ReqRelease struct {
	CurrAbbr string
}

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

// Release returns release by release's ID
func (s *DatabaseService) Release(releaseID int) (*Release, error) {
	params := url.Values{}
	params.Set("CurrAbbr", s.currency)

	var release *Release
	if err := request(s.url+releasesURI+strconv.Itoa(releaseID), params, &release); err != nil {
		return nil, err
	}

	return release, nil
}

// ReleaseRating serves response for community release rating request
type ReleaseRating struct {
	ID     int    `json:"release_id"`
	Rating Rating `json:"rating"`
}

// ReleaseRating retruns community release rating
func (s *DatabaseService) ReleaseRating(releaseID int) (*ReleaseRating, error) {
	var rating *ReleaseRating
	if err := request(s.url+releasesURI+strconv.Itoa(releaseID)+"/rating", nil, &rating); err != nil {
		return nil, err
	}

	return rating, nil
}

// Artist ...
type Artist struct {
	Namevariations []string `json:"namevariations"`
	Profile        string   `json:"profile"`
	ReleasesURL    string   `json:"releases_url"`
	ResourceURL    string   `json:"resource_url"`
	URI            string   `json:"uri"`
	URLs           []string `json:"urls"`
	DataQuality    string   `json:"data_quality"`
	ID             int      `json:"id"`
	Images         []Image  `json:"images"`
	Members        []Member `json:"members"`
}

// Artist represents a person in the discogs database
func (s *DatabaseService) Artist(artistID int) (*Artist, error) {
	var artist *Artist
	if err := request(s.url+artistsURI+strconv.Itoa(artistID), nil, &artist); err != nil {
		return nil, err
	}
	return artist, nil
}

// ArtistReleases ...
type ArtistReleases struct {
	Pagination Page            `json:"pagination"`
	Releases   []ReleaseSource `json:"releases"`
}

// ArtistReleases returns a list of releases and masters associated with the artist.
func (s *DatabaseService) ArtistReleases(artistID int, pagination *Pagination) (*ArtistReleases, error) {
	var releases *ArtistReleases
	if err := request(s.url+artistsURI+strconv.Itoa(artistID)+"/releases", pagination.toParams(), &releases); err != nil {
		return nil, err
	}
	return releases, nil
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
func (s *DatabaseService) Label(labelID int) (*Label, error) {
	var label *Label
	if err := request(s.url+labelsURI+strconv.Itoa(labelID), nil, &label); err != nil {
		return nil, err
	}
	return label, nil
}

// LabelReleases is a list of Releases associated with the label.
type LabelReleases struct {
	Pagination Page            `json:"pagination"`
	Releases   []ReleaseSource `json:"releases"`
}

// LabelReleases returns a list of Releases associated with the label.
func (s *DatabaseService) LabelReleases(labelID int, pagination *Pagination) (*LabelReleases, error) {
	var releases *LabelReleases
	if err := request(s.url+labelsURI+strconv.Itoa(labelID)+"/releases", pagination.toParams(), &releases); err != nil {
		return nil, err
	}
	return releases, nil
}

// Master resource represents a set of similar releases.
// Masters (also known as `master releases`) have a `main release` which is often the chronologically earliest.
type Master struct {
	Styles         []string `json:"styles"`
	Genres         []string `json:"genres"`
	Videos         []Video  `json:"videos"`
	Title          string   `json:"title"`
	MainRelease    int      `json:"main_release"`
	MainReleaseURL string   `json:"main_release_url"`
	URI            string   `json:"uri"`
	Artists        []Artist `json:"artists"`
	VersionURL     string   `json:"version_url"`
	Year           int      `json:"year"`
	Images         []Image  `json:"images"`
	ResourceURL    string   `json:"resource_url"`
	Tracklist      []Track  `json:"tracklist"`
	ID             int      `json:"id"`
	DataQuality    string   `json:"data_quality"`
}

// Master returns a master release
func (s *DatabaseService) Master(masterID int) (*Master, error) {
	var master *Master
	if err := request(s.url+mastersURI+strconv.Itoa(masterID), nil, &master); err != nil {
		return nil, err
	}
	return master, nil
}

// MasterVersions retrieves a list of all releases that are versions of this master.
type MasterVersions struct {
	Pagination Page      `json:"pagination"`
	Versions   []Version `json:"versions"`
}

// MasterVersions retrieves a list of all Releases that are versions of this master
func (s *DatabaseService) MasterVersions(masterID int, pagination *Pagination) (*MasterVersions, error) {
	var versions *MasterVersions
	if err := request(s.url+mastersURI+strconv.Itoa(masterID)+"/versions", pagination.toParams(), &versions); err != nil {
		return nil, err
	}
	return versions, nil
}
