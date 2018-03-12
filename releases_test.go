package discogs

import (
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

	d := initDiscogsClient(t, &Options{URL: ts.URL})
	release, err := d.Release.Release(8138518)
	if err != nil {
		t.Fatalf("failed to get release: %s", err)
	}

	if release.Title != expectedTitle {
		t.Fatalf("release title got=%s want=%s ", expectedTitle, release.Title)
	}
}
