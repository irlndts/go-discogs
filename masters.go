package discogs

import (
	"github.com/irlndts/go-apirequest"
	"net/http"
)

type MasterService struct {
	api *apirequest.API
}

type MasterParams struct {
	Master_id string
}

type MasterVersionParams struct {
	Master_id string
	Page      int
	Per_page  int
}

type Master struct {
	Styles           []string `json:"styles"`
	Genres           []string `json:"genres"`
	Videos           []Video  `json:"videos"`
	Title            string   `json:"title"`
	Main_release     int      `json:"main_release"`
	Main_release_url string   `json:"main_release_url"`
	Uri              string   `json:"uri"`
	Artists          []Artist `json:"artists"`
	Version_url      string   `json:"version_url"`
	Year             int      `json:"year"`
	Images           []Image  `json:"images"`
	Resource_url     string   `json:"resource_url"`
	Tracklist        []Track  `json:"tracklist"`
	Id               int      `json:"id"`
	Data_quality     string   `json:"data_quality"`
}

type MasterVersions struct {
	Pagination Page      `json:"pagination"`
	Versions   []Version `json:"versions"`
}

func newMasterService(api *apirequest.API) *MasterService {
	return &MasterService{
		api: api.Path("masters/"),
	}
}

func (self *MasterService) Master(params *MasterParams) (*Master, *http.Response, error) {
	master := new(Master)
	apiError := new(APIError)
	resp, err := self.api.New().Get(params.Master_id).Receive(master, apiError)
	return master, resp, relevantError(err, *apiError)
}

func (self *MasterService) Versions(params *MasterVersionParams) (*MasterVersions, *http.Response, error) {
	versions := new(MasterVersions)
	apiError := new(APIError)
	resp, err := self.api.New().Get(params.Master_id+"/versions").QueryStruct(params).Receive(versions, apiError)
	return versions, resp, relevantError(err, *apiError)
}
