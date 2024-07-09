package customerror

import "fmt"

// ValidationError represents the structure of a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

// CustomError is a custom error struct that implements the error interface
type CustomError struct {
	HTTPCode        int               `json:"httpCode"`
	ErrorCode       string            `json:"errorCode"`
	Message         string            `json:"message"`
	ValidationError []ValidationError `json:"validationError,omitempty"`

	StackTrace  any            `json:"-"`
	Placeholder map[string]any `json:"-"`
}

// Error returns the error message
func (ce *CustomError) Error() string {
	return fmt.Sprintf("Err With Code: %d Message: %s", ce.HTTPCode, ce.Message)
}

// WithStackTrace when you need actual error and stack trace
func (ce *CustomError) WithStackTrace(st any) *CustomError {
	ce.StackTrace = st
	return ce
}

func (ce *CustomError) WithPlaceholder(st map[string]any) *CustomError {
	ce.Placeholder = st
	return ce
}
