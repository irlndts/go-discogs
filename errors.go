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
	ErrCurrencyNotSupported = &Error{"currency does not supported"}
	ErrInvalidReleaseID     = &Error{"invalid release id"}
	ErrInvalidSortKey       = &Error{"invalid sort key"}
	ErrInvalidUsername      = &Error{"invalid username"}
	ErrTooManyRequests      = &Error{"too many requests"}
	ErrUnauthorized         = &Error{"authentication required"}
	ErrUserAgentInvalid     = &Error{"invalid user-agent"}
)
