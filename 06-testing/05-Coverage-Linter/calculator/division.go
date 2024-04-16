package calculator

import "errors"

var (
	ErrDivisionIndeterminateForm = errors.New("error: indeterminate form")
	ErrDivisionZeroDenominator = errors.New("error: division by zero")
)

// Division returns the division of two numbers
func Division(a, b int) (result int, err error) {
	// check if both numbers are zero
	if a == 0 && b == 0 {
		err = ErrDivisionIndeterminateForm
		return
	}

	// check if the denominator is zero
	if b == 0 {
		err = ErrDivisionZeroDenominator
		return
	}

	result = a / b
	return
}

func Random(something int) int {
	something++
	return something
}