package rest

import (
	"fmt"
)

// APIError reports a non-2xx response from the T-Invest REST gateway. Reach it
// with errors.As; Body holds the raw error JSON (e.g. {"code":..,"message":..}).
type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf(
		"tinvest rest: status %d, body: %s",
		e.StatusCode,
		e.Body,
	)
}

// GetStatusCode reports the HTTP status (for errors.As probes).
func (e *APIError) GetStatusCode() int { return e.StatusCode }
