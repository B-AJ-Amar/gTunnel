package server

// ValidationError represents a validation error with an HTTP status code
type ValidationError struct {
	Message    string
	StatusCode int
}

func (e *ValidationError) Error() string {
	return e.Message
}
