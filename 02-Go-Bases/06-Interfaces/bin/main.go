package main

import "go-bases/interfaces/lib/logger"

func main() {
	// dependencies
	var lg logger.Logger
	lgTxt := logger.NewTextFile("log.txt")
	lgLocal := logger.NewLocal("red")


	// application
	// ...
	lg = lgTxt
	lg.Info("Hello, world!")

	lg = lgLocal
	lg.Info("Hello, world 2!")

	lg = lgTxt
	lg.Info("Hello, world 3!")
}