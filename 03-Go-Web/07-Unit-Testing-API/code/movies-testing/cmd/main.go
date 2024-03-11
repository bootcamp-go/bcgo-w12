package main

import (
	"fmt"
	"go-web/unit-testing/app/internal/application"
)

func main() {
	// env
	// ...

	// application
	// - config
	app := application.NewDefault(":8080")
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}