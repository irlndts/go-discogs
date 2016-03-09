package discogs

import (
	"fmt"
)

// APIError represents a Discogs API Error response
type APIError struct {
	Message string `json:"message"`
}

func (e APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("discogs: %v", e.Message)
	}
	return ""
}

// Empty returns true if empty. Otherwise, at least 1 error message/code is
// present and false is returned.
func (e APIError) Empty() bool {
	if e.Message == "" {
		return true
	}
	return false
}

// relevantError returns any non-nil http-related error (creating the request,
// getting the response, decoding) if any. If the decoded apiError is non-zero
// the apiError is returned. Otherwise, no errors occurred, returns nil.
func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}
	if apiError.Empty() {
		return nil
	}
	return apiError
}
