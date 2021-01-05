package discogs

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func CollectionServer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	switch r.URL.Path {
	case "/users/" + testUsername + "/collection/folders":
		w.WriteHeader(http.StatusOK)
		if _, err := io.WriteString(w, collectionJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case "/users/" + testUsername + "/collection/folders/0":
		w.WriteHeader(http.StatusOK)
		if _, err := io.WriteString(w, folderJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TestCollectionServiceFolder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(CollectionServer))
	defer ts.Close()

	d := initDiscogsClient(t, &Options{URL: ts.URL, Username: testUsername})

	folder, err := d.Folder(0)
	if err != nil {
		t.Fatalf("failed to get folder: %s", err)
	}

	json, err := json.Marshal(folder)
	if err != nil {
		t.Fatalf("failed to marshal folder: %s", err)
	}

	compareJson(t, string(json), folderJson)
}

func TestCollectionServiceCollection(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(CollectionServer))
	defer ts.Close()

	d := initDiscogsClient(t, &Options{URL: ts.URL, Username: testUsername})

	collection, err := d.CollectionFolders()
	if err != nil {
		t.Fatalf("failed to get collection: %s", err)
	}

	json, err := json.Marshal(collection)
	if err != nil {
		t.Fatalf("failed to marshal collection: %s", err)
	}

	compareJson(t, string(json), collectionJson)
}
