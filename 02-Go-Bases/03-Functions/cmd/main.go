package main

import "go-bases/functions/internal/calculator"

func main() {
	// numbers
	// a := 10
	// b := 5

	// division
	// result, err := calculator.Divide(a, b)
	// _____________
	// bad practice
	// if err == "" {
	// 	result = calculator.Add(result, 1)

	// 	result, err = calculator.Divide(result, 2)
	// 	if err == "" {
	// 		result = calculator.Multiply(result, 2)
	// 	}
	// }

	// result, err := calculator.Divide(a, b)
	// if err != "" {
	// 	println(err)
	// 	return
	// }

	// result = calculator.Add(result, 1)
	// println(result)

	// numbers
	a, b := 10, 5

	// division
	fnDiv, err := calculator.Orchestrator(calculator.OperatorDiv)
	if err != "" {
		println(err)
		return
	}

	result, err := fnDiv(a, b)
	if err != "" {
		println(err)
		return
	}

	println(result)
}