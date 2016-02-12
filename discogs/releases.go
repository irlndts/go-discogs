package discogs

import (
	"github.com/irlndts/go-apirequest"
)

type ReleaseService struct {
	api *apirequest.API
}

func newReleaseService(api *apirequest.API) *ReleaseService {
	return &ReleaseService{
		api: api.Path("releases/"),
	}
}
