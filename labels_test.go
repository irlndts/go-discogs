package discogs

import (
	"fmt"
	"net/http"
	"testing"
)

func TestLabelService_Label(t *testing.T) {
	expectedId := 1000

	client := &http.Client{}
	d := NewClient(client).UserAgent("UnitTestClient/0.0.1 +https://github.com/irlndts/go-discogs")
	label, _, err := d.Label.Label(&LabelParams{Label_id: "1000"})

	check(t, err)
	assert(t, label.Id == expectedId, fmt.Sprintf("Release.Title looked for %s, and received %s ", expectedId, label.Id))
}
