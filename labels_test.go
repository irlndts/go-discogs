package discogs

import (
	"fmt"
	"testing"
)

func TestLabelService_Label(t *testing.T) {
	expectedId := 1000

	d := NewClient(testUserAgent, testToken)
	label, _, err := d.Label.Label(&LabelParams{Label_id: "1000"})

	check(t, err)
	assert(t, label.Id == expectedId, fmt.Sprintf("Release.Title looked for %s, and received %s ", expectedId, label.Id))
}

func TestLabelService_Releases(t *testing.T) {
	expectedId := "Cha Cha Twist"

	d := NewClient(testUserAgent, testToken)
	label, _, err := d.Label.Releases(&LabelParams{Label_id: "1000"})

	check(t, err)
	assert(t, label.Releases[0].Title == expectedId, fmt.Sprintf("Release.Title looked for %s, and received %s ", expectedId, label.Releases[0].Title))
}
