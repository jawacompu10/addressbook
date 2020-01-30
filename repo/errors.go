package repo

import (
	"fmt"
	"net/http"
)

// InvalidIDError is the error returned when an incoming request sends an invalid ID
type InvalidIDError struct {
	ID string
}

// GetStatusCode returns the status code this error should return to a HTTP client
func (ie *InvalidIDError) GetStatusCode() int {
	return http.StatusBadRequest
}

// Error method makes this type an error,
// and this is the error message that will be printed
func (ie *InvalidIDError) Error() string {
	return fmt.Sprintf("Invalid ID: %s", ie.ID)
}
