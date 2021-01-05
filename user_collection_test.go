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

	case "/users/" + testUsername + "/collection/folders/0/releases":
		w.WriteHeader(http.StatusOK)
		if _, err := io.WriteString(w, collectionItemsByFolderJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case "/users/" + testUsername + "/collection/releases/12934893":
		w.WriteHeader(http.StatusOK)
		if _, err := io.WriteString(w, collectionItemsByRelease); err != nil {
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

	d := initDiscogsClient(t, &Options{URL: ts.URL})

	folder, err := d.Folder(testUsername, 0)
	if err != nil {
		t.Fatalf("failed to get folder: %s", err)
	}

	json, err := json.Marshal(folder)
	if err != nil {
		t.Fatalf("failed to marshal folder: %s", err)
	}

	compareJson(t, string(json), folderJson)
}

func TestCollectionServiceCollectionFolders(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(CollectionServer))
	defer ts.Close()

	d := initDiscogsClient(t, &Options{URL: ts.URL})

	collection, err := d.CollectionFolders(testUsername)
	if err != nil {
		t.Fatalf("failed to get collection: %s", err)
	}

	json, err := json.Marshal(collection)
	if err != nil {
		t.Fatalf("failed to marshal collection: %s", err)
	}

	compareJson(t, string(json), collectionJson)
}

func TestCollectionServiceCollectionItemsByFolder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(CollectionServer))
	defer ts.Close()

	d := initDiscogsClient(t, &Options{URL: ts.URL})

	items, err := d.CollectionItemsByFolder(testUsername, 0, &Pagination{Sort: "artist", SortOrder: "desc", PerPage: 2})

	if err != nil {
		t.Fatalf("failed to get collection items: %s", err)
	}

	json, err := json.Marshal(items)
	if err != nil {
		t.Fatalf("failed to marshal collection items: %s", err)
	}

	compareJson(t, string(json), collectionItemsByFolderJson)
}

func TestCollectionServiceCollectionItemsByRelease(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(CollectionServer))
	defer ts.Close()

	d := initDiscogsClient(t, &Options{URL: ts.URL})

	items, err := d.CollectionItemsByRelease(testUsername, 12934893, &Pagination{PerPage: 2})

	if err != nil {
		t.Fatalf("failed to get collection items: %s", err)
	}

	json, err := json.Marshal(items)
	if err != nil {
		t.Fatalf("failed to marshal collection items: %s", err)
	}

	compareJson(t, string(json), collectionItemsByRelease)
}
