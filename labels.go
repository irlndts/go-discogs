package discogs

import (
	"github.com/irlndts/go-apirequest"
	"net/http"
)

type LabelService struct {
	api *apirequest.API
}

type LabelParams struct {
	Label_id string
}

type Label struct {
	Profile      string     `json:"profile"`
	Releases_url string     `json:"releases_url"`
	Name         string     `json:"name"`
	Contact_info string     `json:"contact_info"`
	Uri          string     `json:"uri"`
	Sublabels    []Sublable `json:"sublabels"`
	Urls         []string   `json:"urls"`
	Images       []Image    `json:"images"`
	Resource_url string     `json:"resource_url"`
	Id           int        `json:"id"`
	Data_quality string     `json:"data_quality"`
}

func newLabelService(api *apirequest.API) *LabelService {
	return &LabelService{
		api: api.Path("labels/"),
	}
}

func (self *LabelService) Label(params *LabelParams) (*Label, *http.Response, error) {
	label := new(Label)
	apiError := new(APIError)

	resp, err := self.api.New().Get(params.Label_id).Receive(label, apiError)
	return label, resp, relevantError(err, *apiError)
}
