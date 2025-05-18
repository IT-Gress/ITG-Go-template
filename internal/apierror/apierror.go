package apierror

// APIError provides a custom error type for the application.
type APIError struct {
	Code    int    // HTTP status code
	Message string // user-facing message
	Err     error  // wrapped base error
}

// Error implements the error interface.
func (e *APIError) Error() string {
	return e.Message + ": " + e.Err.Error()
}

// Unwrap returns the wrapped error.
func (e *APIError) Unwrap() error {
	return e.Err
}

// New creates a new instance of APIError.
func New(code int, msg string, err error) *APIError {
	return &APIError{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}
