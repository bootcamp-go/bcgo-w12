package main

import "go-web/middlewares/app/internal/application"

func main() {
	// env
	// ...

	// application
	// - config
	app := application.NewServerChi("localhost", "8080")
	// - run
	if err := app.Run(); err != nil {
		println(err.Error())
		return
	}
}