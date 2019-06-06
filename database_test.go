package discogs

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ReleaseServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := io.WriteString(w, `{"title":"Elephant Riddim"}`); err != nil {
		panic(err)
	}
}

func TestReleaseServiceRelease(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ReleaseServer))
	defer ts.Close()

	expectedTitle := "Elephant Riddim"
	d := initDiscogsClient(t, &Options{URL: ts.URL})
	release, err := d.Database.Release(8138518)
	if err != nil {
		t.Fatalf("failed to get release: %s", err)
	}

	if release.Title != expectedTitle {
		t.Fatalf("release title got=%s want=%s ", expectedTitle, release.Title)
	}
}
