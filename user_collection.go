package discogs

import (
	"net/url"
	"strconv"
)

// CollectionService is an interface to work with collection.
type CollectionService interface {
	// Retrieve a list of folders in a user’s collection.
	CollectionFolders() (*CollectionFolders, error)
	// Retrieve metadata about a folder in a user’s collection
	Folder(folderID int) (*Folder, error)
}

type collectionService struct {
	url      string
	username string
}

func newCollectionService(url string, username string) CollectionService {
	return &collectionService{
		url:      url,
		username: username,
	}
}

// Folder serves folder response from discogs.
type Folder struct {
	ID          int    `json:"id"`
	Count       int    `json:"count"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

func (s *collectionService) Folder(folderID int) (*Folder, error) {
	params := url.Values{}
	var folder *Folder
	err := request(s.url+"/"+s.username+"/collection/folders/"+strconv.Itoa(folderID), params, &folder)
	return folder, err
}

// Collection serves collection response from discogs.
type CollectionFolders struct {
	Folders []Folder `json:"folders"`
}

func (s *collectionService) CollectionFolders() (*CollectionFolders, error) {
	params := url.Values{}
	var collection *CollectionFolders
	err := request(s.url+"/"+s.username+"/collection/folders", params, &collection)
	return collection, err
}
