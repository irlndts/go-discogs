package discogs

import "strconv"

// ArtistService ...
type ArtistService struct {
	url string
}

func newArtistService(url string) *ArtistService {
	return &ArtistService{
		url: url,
	}
}

// Artist ...
type Artist struct {
	Namevariations []string `json:"namevariations"`
	Profile        string   `json:"profile"`
	Releases_url   string   `json:"releases_url"`
	Resource_url   string   `json:"resource_url"`
	Uri            string   `json:"uri"`
	Urls           []string `json:"urls"`
	Data_quality   string   `json:"data_quality"`
	Id             int      `json:"id"`
	Images         []Image  `json:"images"`
	Members        []Member `json:"members"`
}

// ArtistReleases ...
type ArtistReleases struct {
	Pagination Page            `json:"pagination"`
	Releases   []ReleaseSource `json:"releases"`
}

// Artist represents a person in the discogs database
func (s *ArtistService) Artist(artistID int) (*Artist, error) {
	var artist *Artist
	if err := request(s.url+strconv.Itoa(artistID), nil, &artist); err != nil {
		return nil, err
	}
	return artist, nil
}

// Releases returns a list of releases and masters associated with the artist.
func (s *ArtistService) Releases(artistID int, pagination *Pagination) (*ArtistReleases, error) {
	var releases *ArtistReleases
	if err := request(s.url+strconv.Itoa(artistID)+"/releases", pagination.toParams(), &releases); err != nil {
		return nil, err
	}
	return releases, nil
}
