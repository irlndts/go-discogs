package discogs

import (
	"fmt"
	"strings"
)

// Error represents a Discogs API error
type Error struct {
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("discogs error: %s", strings.ToLower(e.Message))
}

// APIErrors
var (
	ErrUnauthorized         = &Error{"authentication required"}
	ErrCurrencyNotSupported = &Error{"currency does not supported"}
	ErrUserAgentInvalid     = &Error{"invalid user-agent"}
	ErrInvalidReleaseID     = &Error{"invalid release id"}
	ErrInvalidSortKey       = &Error{"invalid sort key"}
	ErrInvalidUsername      = &Error{"invalid username"}
)
