package calculator

var (
	// ErrMsgDivisionIndeterminate is the error message when numerator and denominator are both zero
	ErrMsgDivisionIndeterminate = "error: division indeterminate"
	// ErrMsgDivisionByZero is the error message when the denominator is zero
	ErrMsgDivisionByZero = "error: division by zero"
)

type MathOperator int

const (
	OperatorSum MathOperator = iota
	OperatorSub
	OperatorMul
	OperatorDiv
)

type MathOperation func(int, int) (int, string)

// Add calculates the sum of two numbers
// this numbers are integers
//
//	func Add(a, b int) int {
//		return a + b
//	}
// func Add(a, b int) int {
// 	var result int
// 	result = a + b

//		return result
//	}
func Add(a, b int) (result int, err string) {
	result = a + b
	return
}

// Subtract calculates the difference of two numbers
func Subtract(a, b int) (result int, err string) {
	result = a - b
	return
}

// Multiply calculates the product of two numbers
func Multiply(a, b int) (result int, err string) {
	result = a * b
	return
}

// Divide calculates the division of two numbers
// if the denominator is zero, it returns an error message
func Divide(a, b int) (result int, err string) {
	// check if both numerator and denominator are zero
	if a == 0 && b == 0 {
		err = ErrMsgDivisionIndeterminate
		return
	}

	// check if the denominator is zero
	if b == 0 {
		err = ErrMsgDivisionByZero
		return
	}

	result = a / b
	return
}