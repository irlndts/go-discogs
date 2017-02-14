package discogs

import (
	"testing"
)

const (
	testUserAgent = "UnitTestClient/0.0.2 +https://github.com/irlndts/go-discogs"
	testToken     = ""
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
