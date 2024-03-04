package main

import (
	"errors"
	"go-bases/error-handling/app-02/internal/calculator"
)

func main() {
	// numbers
	a, b := 0, 0

	// division
	result, err := calculator.Divide(a, b)
	if err != nil {
		var errMath *calculator.MathError  // errMath -> 0x000a | nil

		// ChangeSomething(&errMath)

		if errors.As(err, &errMath) {      // errors.As(err, errMath): e -> 0x000b | nil
			switch errMath.Code {
			case 1000:
				println("Indeterminate division")
			case 1001:
				println("Division by zero")
			default:
				println("A math error occurred:", err.Error())
			}
		
			println("full error:", err.Error())
			return
		}

		// // unwrap the error // not recommended, unless specific design
		// err = errors.Unwrap(err)

		// check if the error is a MathError
		// errMath, ok := err.(*calculator.MathError)
		// if ok {
		// 	switch errMath.Code {
		// 	case 1000:
		// 		println("Indeterminate division")
		// 	case 1001:
		// 		println("Division by zero")
		// 	default:
		// 		println("A math error occurred:", err.Error())
		// 	}
		// 	return
		// }

		println("Unknown error:", err.Error())
		return
	}

	// result
	println("The result of the division is", result)
}

// func ChangeSomething(err **calculator.MathError) {
// 	// err -> 0x000b | nil - 0x0
// 	(*err) = &calculator.MathError{
// 		Code: 1000,
// 	}

// 	// err -> 0x000b | 0x000c

// }