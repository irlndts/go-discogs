package discogs

import (
	"testing"
)

func check(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
	}
}

func assert(t *testing.T, condition bool, assertion string) {
	if !condition {
		t.Errorf("Assertion failed: %v", assertion)
	}
}
