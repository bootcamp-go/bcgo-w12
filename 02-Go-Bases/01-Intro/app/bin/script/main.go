package main

import (
	"go-bases/intro/app/lib/calculator"

	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Red("Hello, World!")
	fmt.Println(calculator.Value)

	result := calculator.Sum(1, 2)

	fmt.Println(result)
}