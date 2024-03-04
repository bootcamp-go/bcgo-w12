package calculator

import "fmt"

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
		err = New(1000, a, b)
		err = fmt.Errorf("%w. %s", err, "Indeterminate division")
		return
	}

	// check if the denominator is zero
	if b == 0 {
		err = New(1001, a, b)
		err = fmt.Errorf("%w. %s", err, "Division by zero")
		return
	}

	result = a / b
	return
}