package discogs

import (
	"net/http"

	"github.com/irlndts/go-apirequest"
)

type SearchRequest struct {
	Release_title string
	Artist        string

	Page    int
	PerPage int
}

type Search struct {
	Pagination Page          `json:"pagination"`
	Results    []interface{} `json:"results"`
}

type SearchService struct {
	api *apirequest.API
}

func newSearchService(api *apirequest.API) *SearchService {
	return &SearchService{
		api: api.Path("database/search"),
	}
}

func (self *SearchService) Search(params *SearchRequest) (*Search, *http.Response, error) {
	search := new(Search)
	apiError := new(APIError)

	resp, err := self.api.QueryStruct(params).Receive(search, apiError)
	return search, resp, relevantError(err, *apiError)
}
