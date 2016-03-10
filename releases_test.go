package discogs

import (
	"fmt"
	"net/http"
	"testing"
)

func TestReleaseService_Release(t *testing.T) {
	expectedTitle := "Elephant Riddim"

	client := &http.Client{}
	d := NewClient(client).UserAgent("UnitTestClient/0.0.1 +https://github.com/irlndts/go-discogs")
	release, _, err := d.Release.Release(&ReleaseParams{Release_id: "8138518"})

	check(t, err)
	assert(t, release.Title == expectedTitle, fmt.Sprintf("Release.Title looked for %s, and received %s ", expectedTitle, release.Title))
}
