package discogs

import (
	"fmt"
	"testing"
)

func TestReleaseService_Release(t *testing.T) {
	expectedTitle := "Elephant Riddim"

	d := NewClient(testUserAgent, testToken)
	release, _, err := d.Release.Release(&ReleaseParams{Release_id: "8138518"})

	check(t, err)
	assert(t, release.Title == expectedTitle, fmt.Sprintf("Release.Title looked for %s, and received %s ", expectedTitle, release.Title))
}
