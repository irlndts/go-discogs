package discogs

import (
	"github.com/irlndts/go-apirequest"
	"net/http"
)

type ReleaseService struct {
	api *apirequest.API
}

type ReleaseParams struct {
	Release_id string
}

type Release struct {
	Title              string        `json:"title"`
	Id                 int           `json:"id"`
	Artists            []artist      `json:"artists"`
	Data_quality       string        `json:"data_quality"`
	Thumb              string        `json:"thumb"`
	Community          community     `json:"community"`
	Companies          []company     `json:"companies"`
	Country            string        `json:"country"`
	Date_added         string        `json:"date_added"`
	Date_changed       string        `json:"date_changed"`
	Estimated_weight   int           `json:"estimated_weight"`
	Extraartists       []extraartist `json:"extraartists"`
	Format_quantity    int           `json:"format_quantity"`
	Formats            []format      `json:"formats"`
	Genres             []string      `json:"genres"`
	Identifiers        []identifier  `json:"identifiers"`
	Images             []image       `json:"images"`
	Labels             []label       `json:"labels"`
	Master_id          int           `json:"master_id"`
	Master_url         string        `json:"master_url"`
	Notes              string        `josn:"notes"`
	Released           string        `json:"released"`
	Released_formatted string        `json:"released_formatted"`
	Resource_url       string        `json:"resource_url"`
	Status             string        `json:"status"`
	Styles             []string      `json:"styles"`
	Tracklist          []tracklist   `json:"tracklist"`
	Uri                string        `json:"uri"`
	Videos             []video       `json:"videos"`
	Year               int           `json:"year"`
}

type video struct {
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Embed       bool   `json:"embed"`
	Title       string `json:"title"`
	Uri         string `json:"uri"`
}

type tracklist struct {
	Duration string `json:"duration"`
	Position string `json:"position"`
	Title    string `json:"title"`
	Type     string `json:"type_"`
}

type label struct {
	Catno        string `json:"catno"`
	Entity_type  string `json:"entity_type"`
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Resource_url string `json:"resource_url"`
}

type image struct {
	Height       int    `json:"height"`
	Width        int    `json:"width"`
	Resource_url string `json:"resource_url"`
	Type         string `json:"type"`
	Uri          string `json:"uri"`
	Uri150       string `json:"uri150"`
}

type identifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type format struct {
	Descriptions []string `json:"descriptions"`
	Name         string   `json:"name"`
	Qty          string   `json:"qty"`
}

type extraartist struct {
	Anv          string `json:"anv"`
	Id           int    `json:"id"`
	Join         string `json:"join"`
	Name         string `json:"name"`
	Resource_url string `json:"resource_url"`
	Role         string `json:"role"`
	Tracks       string `json:"tracks"`
}

type company struct {
	Catno            string `json:"catno"`
	Entity_type      string `json:"entity_type"`
	Entity_type_name string `json:"entity_type_name"`
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Resource_url     string `json:"resource_url"`
}

type community struct {
	Contributors []contributor `json:"contributors"`
	Data_quality string        `json:"string"`
	Have         int           `json:"have"`
	Rating       rating        `json:"rating"`
	Status       string        `json:"status"`
	Submitter    submitter     `json:"submitter"`
	Want         int           `json:"want"`
}

type submitter struct {
	Resource_url string `json:"resource_url"`
	Username     string `json:"username"`
}

type rating struct {
	Average float32 `json:"average"`
	Count   int     `json:"count"`
}

type contributor struct {
	Resource_url string `json:"resource_url"`
	Username     string `json:"username"`
}

type artist struct {
	Anv          string `json:"anv"`
	Id           int    `json:"id"`
	Join         string `json:"join"`
	Name         string `json:"name:`
	Resource_url string `json:"resource_url"`
	Role         string `json:"role"`
	Tracks       string `json:"tracks"`
}

func newReleaseService(api *apirequest.API) *ReleaseService {
	return &ReleaseService{
		api: api.Path("releases/"),
	}
}

func (self *ReleaseService) Release(params *ReleaseParams) (*Release, *http.Response, error) {
	release := new(Release)
	apiError := new(APIError)
	resp, err := self.api.New().Get(params.Release_id).QueryStruct(params).Receive(release, apiError)
	return release, resp, relevantError(err, *apiError)
}
