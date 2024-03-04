package calculator

import "fmt"

// New returns a new MathError
func New(code int, args ...any) error {
	return &MathError{Code: code, Args: args}
}

// MathError is the error type for math operations
type MathError struct {
	// Code is the error code of all math errors
	Code int
	// Args is the arguments used in the math operation
	Args []any
}

// Error returns the error message
func (e *MathError) Error() string {
	return fmt.Sprintf("Math error: code - %d, args - %v", e.Code, e.Args)
}