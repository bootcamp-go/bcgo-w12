package tools

import (
	"fmt"
)

type FieldError struct {
	Field string
	Msg   string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// CheckFieldExistance is a function that checks if the required fields exist in the fields map
func CheckFieldExistance(fields map[string]any, requiredFields ...string) (err error) {
	for _, field := range requiredFields {
		if _, ok := fields[field]; !ok {
			err = &FieldError{
				Field: field,
				Msg:   "field is required",
			}
			return
		}
	}
	return
}