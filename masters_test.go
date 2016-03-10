package discogs

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMasterService_Master(t *testing.T) {
	expectedTitle := "Elephant Riddim"

	client := &http.Client{}
	d := NewClient(client).UserAgent("UnitTestClient/0.0.1 +https://github.com/irlndts/go-discogs")
	master, _, err := d.Master.Master(&MasterParams{Master_id: "960657"})

	check(t, err)
	assert(t, master.Title == expectedTitle, fmt.Sprintf("master.Title looked for %s, and received %s ", expectedTitle, master.Title))
}

func TestMasterService_Versions(t *testing.T) {
	expectedTitle := "Stardiver"

	client := &http.Client{}
	d := NewClient(client).UserAgent("UnitTestClient/0.0.1 +https://github.com/irlndts/go-discogs")
	versions, _, err := d.Master.Versions(&MasterVersionParams{Master_id: "1000", Page: 1, Per_page: 1})

	check(t, err)
	assert(t, versions.Versions[0].Title == expectedTitle, fmt.Sprintf("master.Title looked for %s, and received %s ", expectedTitle, versions.Versions[0].Title))
}
