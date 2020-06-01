package discogs

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func DatabaseServer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	switch r.URL.Path {
	case "/releases/8138518":
		w.WriteHeader(http.StatusOK)
		if _, err := io.WriteString(w, releaseJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "/masters/718441":
		w.WriteHeader(http.StatusOK)
		if _, err := io.WriteString(w, masterJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "/artists/38661":
		w.WriteHeader(http.StatusOK)
		if _, err := io.WriteString(w, artistJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func compareJson(t *testing.T, got, want string) {
	var g, w interface{}
	if err := json.Unmarshal([]byte(got), &g); err != nil {
		log.Fatalf("failed to unmarshal json: %s", err)
	}
	if err := json.Unmarshal([]byte(want), &w); err != nil {
		log.Fatalf("failed to unmarshal json: %s", err)
	}

	if diff := cmp.Diff(g, w); diff != "" {
		t.Errorf("(-want +got)\n%s", diff)
	}
}

func TestDatabaseServiceRelease(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(DatabaseServer))
	defer ts.Close()

	d := initDiscogsClient(t, &Options{URL: ts.URL})
	release, err := d.Release(8138518)
	if err != nil {
		t.Fatalf("failed to get release: %s", err)
	}

	json, err := json.Marshal(release)
	if err != nil {
		t.Fatalf("failed to marshal release: %s", err)
	}

	compareJson(t, string(json), releaseJson)
}

func TestDatabaseServiceMaster(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(DatabaseServer))
	defer ts.Close()

	d := initDiscogsClient(t, &Options{URL: ts.URL})
	master, err := d.Master(718441)
	if err != nil {
		t.Fatalf("failed to get master: %s", err)
	}

	json, err := json.Marshal(master)
	if err != nil {
		t.Fatalf("failed to marshal release: %s", err)
	}
	compareJson(t, string(json), masterJson)
}

func TestDatabaseServiceArtist(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(DatabaseServer))
	defer ts.Close()

	d := initDiscogsClient(t, &Options{URL: ts.URL})
	artist, err := d.Artist(38661)
	if err != nil {
		t.Fatalf("failed to get master: %s", err)
	}

	json, err := json.Marshal(artist)
	if err != nil {
		t.Fatalf("failed to marshal artist: %s", err)
	}
	compareJson(t, string(json), artistJson)
}
