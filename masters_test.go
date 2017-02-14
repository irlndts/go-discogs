package discogs

import (
	"fmt"
	"testing"
)

func TestMasterService_Master(t *testing.T) {
	expectedTitle := "Elephant Riddim"

	d := NewClient(testUserAgent, testToken)
	master, _, err := d.Master.Master(&MasterParams{Master_id: "960657"})

	check(t, err)
	assert(t, master.Title == expectedTitle, fmt.Sprintf("master.Title looked for %s, and received %s ", expectedTitle, master.Title))
}

func TestMasterService_Versions(t *testing.T) {
	expectedTitle := "Stardiver"

	d := NewClient(testUserAgent, testToken)
	versions, _, err := d.Master.Versions(&MasterVersionParams{Master_id: "1000", Page: 1, Per_page: 1})

	check(t, err)
	assert(t, versions.Versions[0].Title == expectedTitle, fmt.Sprintf("master.Title looked for %s, and received %s ", expectedTitle, versions.Versions[0].Title))
}
