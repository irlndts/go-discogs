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
