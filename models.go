package discogs

import (
	"net/url"
	"strconv"
)

// Video ...
type Video struct {
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Embed       bool   `json:"embed"`
	Title       string `json:"title"`
	URI         string `json:"uri"`
}

// Series ...
type Series struct {
	Catno          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ResourceURL    string `json:"resource_url"`
	ThumbnailURL   string `json:"thumbnail_url,omitempty"`
}

// ArtistSource ...
type ArtistSource struct {
	Anv         string `json:"anv"`
	ID          int    `json:"id"`
	Join        string `json:"join"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
	Role        string `json:"role"`
	Tracks      string `json:"tracks"`
}

// Image ...
type Image struct {
	Height      int    `json:"height"`
	Width       int    `json:"width"`
	ResourceURL string `json:"resource_url"`
	Type        string `json:"type"`
	URI         string `json:"uri"`
	URI150      string `json:"uri150"`
}

// Track ...
type Track struct {
	Duration     string         `json:"duration"`
	Position     string         `json:"position"`
	Title        string         `json:"title"`
	Type         string         `json:"type_"`
	Extraartists []ArtistSource `json:"extraartists,omitempty"`
	Artists      []ArtistSource `json:"artists,omitempty"`
}

// LabelSource ...
type LabelSource struct {
	Catno          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ResourceURL    string `json:"resource_url"`
}

// Identifier ...
type Identifier struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
	Value       string `json:"value"`
}

// Format ...
type Format struct {
	Descriptions []string `json:"descriptions"`
	Name         string   `json:"name"`
	Qty          string   `json:"qty"`
}

// Company ...
type Company struct {
	Catno          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ResourceURL    string `json:"resource_url"`
}

// Community ...
type Community struct {
	Contributors []Contributor `json:"contributors"`
	DataQuality  string        `json:"data_quality"`
	Have         int           `json:"have"`
	Rating       Rating        `json:"rating"`
	Status       string        `json:"status"`
	Submitter    Submitter     `json:"submitter"`
	Want         int           `json:"want"`
}

// Submitter ...
type Submitter struct {
	ResourceURL string `json:"resource_url"`
	Username    string `json:"username"`
}

// Rating ...
type Rating struct {
	Average float32 `json:"average"`
	Count   int     `json:"count"`
}

// Contributor ...
type Contributor struct {
	ResourceURL string `json:"resource_url"`
	Username    string `json:"username"`
}

// Page ...
type Page struct {
	PerPage int      `json:"per_page"`
	Items   int      `json:"items"`
	Page    int      `json:"page"`
	URLs    URLsList `json:"urls"`
	Pages   int      `json:"pages"`
}

// URLsList ...
type URLsList struct {
	Last string `json:"last"`
	Next string `json:"next"`
}

// Version ...
type Version struct {
	Catno       string `json:"catno"`
	Country     string `json:"country"`
	Format      string `json:"format"`
	ID          int    `json:"id"`
	Label       string `json:"label"`
	Released    string `json:"released"`
	ResourceURL string `json:"resource_url"`
	Status      string `json:"status"`
	Thumb       string `json:"thumb"`
	Title       string `json:"title"`
}

// Member ...
type Member struct {
	Active      bool   `json:"active"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

// Alias ...
type Alias struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

// Sublable ...
type Sublable struct {
	ResourceURL string `json:"url"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
}

// ReleaseSource ...
type ReleaseSource struct {
	Artist      string `json:"artist"`
	Catno       string `json:"catno"`
	Format      string `json:"format"`
	ID          int    `json:"id"`
	ResourceURL string `json:"resource_url"`
	Status      string `json:"status"`
	Thumb       string `json:"thumb"`
	Title       string `json:"title"`
	Year        int    `json:"year"`
	MainRelease int    `json:"main_release"`
	Role        string `json:"role"`
	Type        string `json:"type"`
}

// Pagination ...
type Pagination struct {
	Sort      string // year, title, format
	SortOrder string // asc, desc
	Page      int
	PerPage   int
}

// toParams converts pagaination params to request values
func (p *Pagination) params() url.Values {
	if p == nil {
		return nil
	}

	params := url.Values{}
	params.Set("sort", p.Sort)
	params.Set("sort_order", p.SortOrder)
	params.Set("page", strconv.Itoa(p.Page))
	params.Set("per_page", strconv.Itoa(p.PerPage))
	return params
}
