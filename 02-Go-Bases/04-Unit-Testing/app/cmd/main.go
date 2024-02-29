package main

import "go-bases/unit-testing/internal/calculator"

func main() {
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