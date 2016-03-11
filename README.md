# REST API 2.0 Discogs.com client

go-discogs is a Go client library for the [Discogs API](https://www.discogs.com/developers/). Check the usage section or try the examples to see how to access the Discogs API.

### Feauteres
 * Database
  * Releases
  * Master Releases
  * Release Versions
  * Artists
  * Artist Releases
  * Label
  * All Label Releases
 
#### ToDo
- Search

 
Install
--------
    go get github.com/irlndts/go-discogs

Usage
---------
The discogs package provides a client for accessing the Discogs API. 
First of all import library and init client variable. According to discogs api documentation you [must provide your user-agent](https://www.discogs.com/developers/#page:home,header:home-general-information). 
```go
  package main
    
  import (
    "github.com/irlndts/go-discogs"
  )
 ```
 ```go
 client := discogs.NewClient().UserAgent("TestDiscogsClient/0.0.1 +example.com")
``` 

#### Releases
```go
  params := &discogs.ReleaseParams{Release_id: "8138518"}
  release, _, err := d.Release.Release(params)
  
  fmt.Println(fmt.Println(release.Artists[0].Name, " - ", release.Title)) // St. Petersburg Ska-Jazz Review  -  Elephant Riddim
```

#### Artists
```go
  params := &discogs.LabelParams{Label_id: "890477", Page: 2, Per_page: 3}
  label, _, err := d.Label.Releases(params)

  for _, release := range label.Releases {
    fmt.Println(release.Title)
  }

  /*
    Someday / I Hate Everything About You
    Spy Potion
    Surf Attack From Russia
  */
```

etc. 
More examples - soon
