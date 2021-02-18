package discogs

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

const testReleaseID = 9893847

func MarketplaceServer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	switch r.URL.Path {
	case "/marketplace" + priceSuggestionsURI + strconv.Itoa(testReleaseID):
		w.WriteHeader(http.StatusOK)
		if _, err := io.WriteString(w, priceSuggestionJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case "/marketplace" + releaseStatsURI + strconv.Itoa(testReleaseID):
		w.WriteHeader(http.StatusOK)
		if _, err := io.WriteString(w, releaseStatsJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TestMarketplacePriceSuggestions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(MarketplaceServer))
	defer ts.Close()

	d := initDiscogsClient(t, &Options{URL: ts.URL})

	suggestion, err := d.PriceSuggestions(testReleaseID)
	if err != nil {
		t.Fatalf("failed to get price suggestion: %s", err)
	}

	json, err := json.Marshal(suggestion)
	if err != nil {
		t.Fatalf("failed to marshal folder: %s", err)
	}

	compareJson(t, string(json), priceSuggestionJson)
}

func TestMarketplaceReleaseStatistics(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(MarketplaceServer))
	defer ts.Close()

	d := initDiscogsClient(t, &Options{URL: ts.URL})

	stats, err := d.ReleaseStatistics(testReleaseID)
	if err != nil {
		t.Fatalf("failed to get price suggestion: %s", err)
	}

	json, err := json.Marshal(stats)
	if err != nil {
		t.Fatalf("failed to marshal folder: %s", err)
	}

	compareJson(t, string(json), releaseStatsJson)
}
