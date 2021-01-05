package discogs

import (
	"encoding/json"
	"fmt"
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

	fmt.Println(r.URL.Path)

	switch r.URL.Path {
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
