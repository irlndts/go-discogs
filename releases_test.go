package discogs

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ReleaseServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"title":"Elephant Riddim"}`)
}

func TestReleaseServiceRelease(t *testing.T) {
	expectedTitle := "Elephant Riddim"

	ts := httptest.NewServer(http.HandlerFunc(ReleaseServer))
	defer ts.Close()

	d, err := NewClient(&Options{URL: ts.URL})
	if err != nil {
		t.Fatalf("failed to create client: %s", err)
	}
	release, err := d.Release.Release(8138518)

	check(t, err)
	assert(t, release.Title == expectedTitle, fmt.Sprintf("Release.Title looked for %s, and received %s ", expectedTitle, release.Title))
}
