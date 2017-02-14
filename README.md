# REST API 2.0 Discogs.com client

go-discogs is a Go client library for the [Discogs API](https://www.discogs.com/developers/). Check the usage section to see how to access the Discogs API.

### Changelog
```
14.02.2017 
- search is implemented
- minor improvements
```

### Feauteres
 * Database
  * [Releases](#releases)
  * Master Releases
  * Release Versions
  * Artists
  * Artist Releases
  * Label
  * All Label Releases
  * [Search](#search)
 
Install
--------
    go get github.com/irlndts/go-discogs

Usage
---------
The discogs package provides a client for accessing the Discogs API. 
First of all import library and init client variable. According to discogs api documentation you [must provide your user-agent](https://www.discogs.com/developers/#page:home,header:home-general-information). 
```go
import "github.com/irlndts/go-discogs"
```
```go
client := discogs.NewClient("TestDiscogsClient/0.0.1 +example.com", "")
``` 
Some requests require authentication (as any user). According to [Discogs](https://www.discogs.com/developers/#page:authentication,header:authentication-discogs-auth-flow), to send requests with Discogs Auth, you have two options: sending your credentials in the query string with key and secret parameters or a token parameter. Key-secret doesn't implemented yet, but token is yes.
```go
client := discogs.NewClient("TestDiscogsClient/0.0.1 +example.com", "sometoken")
``` 

#### Releases
```go
  params := &discogs.ReleaseParams{Release_id: "8138518"}
  release, _, err := client.Release.Release(params)
  
  fmt.Println(fmt.Println(release.Artists[0].Name, " - ", release.Title)) // St. Petersburg Ska-Jazz Review  -  Elephant Riddim
```

#### Artists
```go
  params := &discogs.LabelParams{Label_id: "890477", Page: 2, Per_page: 3}
  label, _, err := client.Label.Releases(params)

  for _, release := range label.Releases {
    fmt.Println(release.Title)
  }

  /*
    Someday / I Hate Everything About You
    Spy Potion
    Surf Attack From Russia
  */
```

#### Search
Issue a search query to discogs database. This endpoint accepts pagination parameters
Authentication (as any user) is required.

Use `SearchRequest` struc to create a request
```go
type SearchRequest struct {
    Q             string // search query (optional)
    Type          string // one of release, master, artist, label (optional)
    Title         string // search by combined “Artist Name - Release Title” title field (optional)
    Release_title string // search release titles (optional)
    Credit        string // search release credits (optional)
    Artist        string // search artist names (optional)
    Anv           string // search artist ANV (optional)
    Label         string // search label names (optional)
    Genre         string // search genres (optional)
    Style         string // search styles (optional)
    Country       string // search release country (optional)
    Year          string // search release year (optional)
    Format        string // search formats (optional)
    Catno         string // search catalog number (optional)
    Barcode       string // search barcodes (optional)
    Track         string // search track titles (optional)
    Submitter     string // search submitter username (optional)
    Contributer   string // search contributor usernames (optional)

    Page     int // optional
    Per_page int // optional
}
```

Example
```go
  request:= &discogs.SearchRequest{Artist: "reggaenauts", Release_title: "river rock", Page: 0, Per_page: 1}
  search, _, err := client.Search(request)

  for _, r := range search.Results {
    fmt.Println(r.Title)
  }
```

etc. 
More examples - soon
