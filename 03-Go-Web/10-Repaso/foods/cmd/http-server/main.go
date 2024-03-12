package main

import (
	"fmt"
	"go-web/repaso/foods/internal/application"
)

func main() {
	// env / config files
	// ...

	// application
	app := application.NewServerChi("localhost:8080")
	if err := app.Start(); err != nil {
		fmt.Println(err)
		return
	}
}