package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	discogsAPI = "https://api.discogs.com/"
)

type Options struct {
	URL       string
	Currency  string
	UserAgent string
	Token     string
}

// Client is a Discogs client for making Discogs API requests.
type Client struct {
	Release *ReleaseService
	Master  *MasterService
	Artist  *ArtistService
	Label   *LabelService
	Search  *SearchService
}

var header *http.Header

// NewClient returns a new Client.
func NewClient(o *Options) (*Client, error) {
	header = &http.Header{}

	if o == nil || o.UserAgent == "" {
		return nil, fmt.Errorf("failed to set user-agent")
	}

	header.Add("User-Agent", o.UserAgent)

	cur, err := currency(o.Currency)
	if err != nil {
		return nil, err
	}

	// set token, it's required for some queries like search
	if o.Token != "" {
		header.Add("Authorization", "Discogs token="+o.Token)
	}

	if o.URL == "" {
		o.URL = discogsAPI
	}

	return &Client{
		Release: newReleaseService(o.URL+"releases/", cur),
		Artist:  newArtistService(o.URL + "artists/"),
		Label:   newLabelService(o.URL + "labels/"),
		Master:  newMasterService(o.URL + "masters/"),
		Search:  newSearchService(o.URL + "database/search"),
	}, nil
}

// currency validates currency for marketplace data.
// Defaults to the authenticated users currency. Must be one of the following:
// USD GBP EUR CAD AUD JPY CHF MXN BRL NZD SEK ZAR
func currency(c string) (string, error) {
	switch c {
	case "USD", "GBP", "EUR", "CAD", "AUD", "JPY", "CHF", "MXN", "BRL", "NZD", "SEK", "ZAR":
		return c, nil
	case "":
		return "USD", nil
	default:
		return "", fmt.Errorf("%v\n", "Invalid currency abbreviation.")
	}
}

func request(path string, params url.Values, resp interface{}) error {
	fmt.Println(path + "?" + params.Encode())
	r, err := http.NewRequest("GET", path+"?"+params.Encode(), nil)
	if err != nil {
		return err
	}
	r.Header = *header

	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &resp)
}
