package discogs

import (
	"fmt"
	"testing"
)

func TestArtistService_Artist(t *testing.T) {
	expectedId := 1000

	d := NewClient(testUserAgent, testToken)
	artist, _, err := d.Artist.Artist(&ArtistParams{Artist_id: "1000"})

	check(t, err)
	assert(t, artist.Id == expectedId, fmt.Sprintf("Release.Title looked for %s, and received %s ", expectedId, artist.Id))
}

func TestArtistService_Releases(t *testing.T) {
	expectedArtist := "Dave Clarke"

	d := NewClient(testUserAgent, testToken)
	releases, _, err := d.Artist.Releases(&ArtistParams{Artist_id: "1000", Sort: "year", Sort_order: "desc"})

	check(t, err)
	assert(t, releases.Releases[0].Artist == expectedArtist, fmt.Sprintf("Releses.Artist looked for %s, and received %s ", expectedArtist, releases.Releases[0].Artist))
}
