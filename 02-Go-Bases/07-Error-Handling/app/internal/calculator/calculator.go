package calculator

import (
	"errors"
	"fmt"
)

var (
	// ErrDivisionIndeterminate is the error message when numerator and denominator are both zero
	ErrDivisionIndeterminate = errors.New("division indeterminate")  // *errors.errorString -> 0x001
	// ErrDivisionByZero is the error message when the denominator is zero
	ErrDivisionByZero = errors.New("division by zero")				 // *errors.errorString -> 0x002
)

type MathOperator int

const (
	OperatorSum MathOperator = iota
	OperatorSub
	OperatorMul
	OperatorDiv
)

type MathOperation func(int, int) (int, error)

// Add calculates the sum of two numbers
func Add(a, b int) (result int, err error) {
	result = a + b
	return
}

// Subtract calculates the difference of two numbers
func Subtract(a, b int) (result int, err error) {
	result = a - b
	return
}

// Multiply calculates the product of two numbers
func Multiply(a, b int) (result int, err error) {
	result = a * b
	return
}

// Divide calculates the division of two numbers
// if the denominator is zero, it returns an error message
func Divide(a, b int) (result int, err error) {
	// check if both numerator and denominator are zero
	if a == 0 && b == 0 {
		// err = ErrDivisionIndeterminate
		err = fmt.Errorf("%w. Numerator: %d, Denominator: %d", ErrDivisionIndeterminate, a, b)
		err = fmt.Errorf("%w. Extra detail", err)
		return
	}

	// check if the denominator is zero
	if b == 0 {
		// err = ErrDivisionByZero
		err = fmt.Errorf("%w. Numerator: %d, Denominator: %d", ErrDivisionByZero, a, b)
		err = fmt.Errorf("%w. Extra detail", err)
		return
	}

	result = a / b
	return
}