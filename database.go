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

// DatabaseService is an interface to work with database.
type DatabaseService interface {
	// Artist represents a person in the discogs database.
	Artist(artistID int) (*Artist, error)
	// ArtistReleases returns a list of releases and masters associated with the artist.
	ArtistReleases(artistID int, pagination *Pagination) (*ArtistReleases, error)
	// Label returns a label.
	Label(labelID int) (*Label, error)
	// LabelReleases returns a list of Releases associated with the label.
	LabelReleases(labelID int, pagination *Pagination) (*LabelReleases, error)
	// Master returns a master release.
	Master(masterID int) (*Master, error)
	// MasterVersions retrieves a list of all Releases that are versions of this master.
	MasterVersions(masterID int, pagination *Pagination) (*MasterVersions, error)
	// Release returns release by release's ID.
	Release(releaseID int) (*Release, error)
	// ReleaseRating retruns community release rating.
	ReleaseRating(releaseID int) (*ReleaseRating, error)
}

type databaseService struct {
	url      string
	currency string
}

func newDatabaseService(url string, currency string) DatabaseService {
	return &databaseService{
		url:      url,
		currency: currency,
	}
}

// Release serves relesase response from discogs.
type Release struct {
	Title             string         `json:"title"`
	ID                int            `json:"id"`
	Artists           []ArtistSource `json:"artists"`
	ArtistsSort       string         `json:"artists_sort"`
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
	Series            []Series       `json:"series"`
	Status            string         `json:"status"`
	Styles            []string       `json:"styles"`
	Tracklist         []Track        `json:"tracklist"`
	URI               string         `json:"uri"`
	Videos            []Video        `json:"videos"`
	Year              int            `json:"year"`
}

func (s *databaseService) Release(releaseID int) (*Release, error) {
	params := url.Values{}
	params.Set("curr_abbr", s.currency)

	var release *Release
	err := request(s.url+releasesURI+strconv.Itoa(releaseID), params, &release)
	return release, err
}

// ReleaseRating serves response for community release rating request.
type ReleaseRating struct {
	ID     int    `json:"release_id"`
	Rating Rating `json:"rating"`
}

func (s *databaseService) ReleaseRating(releaseID int) (*ReleaseRating, error) {
	var rating *ReleaseRating
	err := request(s.url+releasesURI+strconv.Itoa(releaseID)+"/rating", nil, &rating)
	return rating, err
}

// Artist resource represents a person in the Discogs database
// who contributed to a Release in some capacity.
// More information https://www.discogs.com/developers#page:database,header:database-artist
type Artist struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Realname       string   `json:"realname"`
	Members        []Member `json:"members,omitempty"`
	Aliases        []Alias  `json:"aliases,omitempty"`
	Namevariations []string `json:"namevariations"`
	Images         []Image  `json:"images"`
	Profile        string   `json:"profile"`
	ReleasesURL    string   `json:"releases_url"`
	ResourceURL    string   `json:"resource_url"`
	URI            string   `json:"uri"`
	URLs           []string `json:"urls"`
	DataQuality    string   `json:"data_quality"`
}

func (s *databaseService) Artist(artistID int) (*Artist, error) {
	var artist *Artist
	err := request(s.url+artistsURI+strconv.Itoa(artistID), nil, &artist)
	return artist, err
}

// ArtistReleases ...
type ArtistReleases struct {
	Pagination Page            `json:"pagination"`
	Releases   []ReleaseSource `json:"releases"`
}

func (s *databaseService) ArtistReleases(artistID int, pagination *Pagination) (*ArtistReleases, error) {
	var releases *ArtistReleases
	err := request(s.url+artistsURI+strconv.Itoa(artistID)+"/releases", pagination.params(), &releases)
	return releases, err
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

func (s *databaseService) Label(labelID int) (*Label, error) {
	var label *Label
	err := request(s.url+labelsURI+strconv.Itoa(labelID), nil, &label)
	return label, err
}

// LabelReleases is a list of Releases associated with the label.
type LabelReleases struct {
	Pagination Page            `json:"pagination"`
	Releases   []ReleaseSource `json:"releases"`
}

func (s *databaseService) LabelReleases(labelID int, pagination *Pagination) (*LabelReleases, error) {
	var releases *LabelReleases
	err := request(s.url+labelsURI+strconv.Itoa(labelID)+"/releases", pagination.params(), &releases)
	return releases, err
}

// Master resource represents a set of similar releases.
// Masters (also known as `master releases`) have a `main release` which is often the chronologically earliest.
// More information https://www.discogs.com/developers#page:database,header:database-master-release
type Master struct {
	ID                   int            `json:"id"`
	Styles               []string       `json:"styles"`
	Genres               []string       `json:"genres"`
	Title                string         `json:"title"`
	Year                 int            `json:"year"`
	Tracklist            []Track        `json:"tracklist"`
	Notes                string         `json:"notes"`
	Artists              []ArtistSource `json:"artists"`
	Images               []Image        `json:"images"`
	Videos               []Video        `json:"videos"`
	NumForSale           int            `json:"num_for_sale"`
	LowestPrice          float64        `json:"lowest_price"`
	URI                  string         `json:"uri"`
	MainRelease          int            `json:"main_release"`
	MainReleaseURL       string         `json:"main_release_url"`
	MostRecentRelease    int            `json:"most_recent_release"`
	MostRecentReleaseURL string         `json:"most_recent_release_url"`
	VersionsURL          string         `json:"versions_url"`
	ResourceURL          string         `json:"resource_url"`
	DataQuality          string         `json:"data_quality"`
}

func (s *databaseService) Master(masterID int) (*Master, error) {
	var master *Master
	err := request(s.url+mastersURI+strconv.Itoa(masterID), nil, &master)
	return master, err
}

// MasterVersions retrieves a list of all releases that are versions of this master.
type MasterVersions struct {
	Pagination Page      `json:"pagination"`
	Versions   []Version `json:"versions"`
}

func (s *databaseService) MasterVersions(masterID int, pagination *Pagination) (*MasterVersions, error) {
	var versions *MasterVersions
	err := request(s.url+mastersURI+strconv.Itoa(masterID)+"/versions", pagination.params(), &versions)
	return versions, err
}
