package main

import (
	"encoding/json"
	"go-web/env-config/app/internal/application"
	"os"
)

func main() {
	// config (env / config files)
	// - json / yaml / toml
	file, err := os.Open("env.json")
	if err != nil {
		println(err.Error())
		return
	}
	defer file.Close()

	var cfg application.ConfigDefaultApp
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		println(err.Error())
		return
	}

	// application
	// cfg := application.ConfigDefaultApp{
	// 	Title:    os.Getenv("APP_TITLE"),
	// 	Color:    os.Getenv("APP_COLOR"),
	// 	FilePath: os.Getenv("APP_FILE_PATH"),
	// }
	app := application.NewDefaultApp(cfg)

	// run
	if err := app.Run(); err != nil {
		println(err.Error())
		return
	}
}