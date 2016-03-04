package discogs

import (
	"github.com/irlndts/go-apirequest"
	//"io/ioutil"
	"net/http"
)

type ReleaseService struct {
	api *apirequest.API
}

type ReleaseParams struct {
	Release_id int
}

type Release struct {
	Title string `json:"title"`
	//id           int
	//data_quality string
}

func newReleaseService(api *apirequest.API) *ReleaseService {
	return &ReleaseService{
		api: api.Path("releases/"),
	}
}

func (self *ReleaseService) Release(params *ReleaseParams) (*Release, *http.Response, error) {
	release := new(Release)
	resp, err := self.api.New().Get("248504").QueryStruct(params).Receive(release, nil)
	return release, resp, err
}
