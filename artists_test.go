package discogs

import (
	"fmt"
	"net/http"
	"testing"
)

func TestArtistService_Artist(t *testing.T) {
	expectedId := 1000

	client := &http.Client{}
	d := NewClient(client).UserAgent("UnitTestClient/0.0.1 +https://github.com/irlndts/go-discogs")
	artist, _, err := d.Artist.Artist(&ArtistParams{Artist_id: "1000"})

	check(t, err)
	assert(t, artist.Id == expectedId, fmt.Sprintf("Release.Title looked for %s, and received %s ", expectedId, artist.Id))
}

func TestArtistService_Releases(t *testing.T) {
	expectedArtist := "Dave Clarke"

	client := &http.Client{}
	d := NewClient(client).UserAgent("UnitTestClient/0.0.1 +https://github.com/irlndts/go-discogs")
	releases, _, err := d.Artist.Releases(&ArtistParams{Artist_id: "1000", Sort: "year", Sort_order: "desc"})

	check(t, err)
	assert(t, releases.Releases[0].Artist == expectedArtist, fmt.Sprintf("Releses.Artist looked for %s, and received %s ", expectedArtist, releases.Releases[0].Artist))
}
