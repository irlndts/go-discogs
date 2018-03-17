package discogs

import (
	"strconv"
)

// MasterService ...
type MasterService struct {
	url string
}

func newMasterService(url string) *MasterService {
	return &MasterService{
		url: url,
	}
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
func (s *MasterService) Master(masterID int) (*Master, error) {
	var master *Master
	if err := request(s.url+strconv.Itoa(masterID), nil, &master); err != nil {
		return nil, err
	}
	return master, nil
}

// MasterVersions retrieves a list of all releases that are versions of this master.
type MasterVersions struct {
	Pagination Page      `json:"pagination"`
	Versions   []Version `json:"versions"`
}

// Versions retrieves a list of all Releases that are versions of this master
func (s *MasterService) Versions(masterID int, pagination *Pagination) (*MasterVersions, error) {
	var versions *MasterVersions
	if err := request(s.url+strconv.Itoa(masterID)+"/versions", pagination.toParams(), &versions); err != nil {
		return nil, err
	}
	return versions, nil
}
