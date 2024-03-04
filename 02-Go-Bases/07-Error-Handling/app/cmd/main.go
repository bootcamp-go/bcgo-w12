package main

import (
	"errors"
	"go-bases/error-handling/app-01/internal/calculator"
)

func main() {
	// numbers
	a, b := 0, 0

	// division
	result, err := calculator.Divide(a, b)
	if err != nil {
		switch {
		case errors.Is(err, calculator.ErrDivisionIndeterminate):
			println("Math error: division indeterminate")
		case errors.Is(err, calculator.ErrDivisionByZero):
			println("Math error: division by zero")
		default:
			println("Unknown error:", err.Error())
		}
		return

		// switch err {
		// case calculator.ErrDivisionIndeterminate: // *errors.errorString -> 0x001 | no funciona
		// // case errors.New("division indeterminate"): // *errors.errorString -> 0x003 | no funciona
		// 	println("Math error: division indeterminate")
		// case calculator.ErrDivisionByZero:
		// 	println("Math error: division by zero")
		// default:
		// 	println("Unknown error:", err.Error())
		// }
		// return
	}

	// result
	println("The result of the division is", result)
}