package main

import "os"

func main() {
	// Environment variables
	// - get all
	env := os.Environ()
	for _, e := range env {
		println(e)
	}

	println("---")

	// - get specific
	filePath1 := os.Getenv("FILE_PATH_1")
	println("FILE_PATH_1:", filePath1)

	filePath2 := os.Getenv("FILE_PATH_2")
	println("FILE_PATH_2:", filePath2)

	filePath3 := os.Getenv("FILE_PATH_3")
	println("FILE_PATH_3:", filePath3)


	// - modify
	os.Setenv("FILE_PATH", "/tmp/file.txt")

	// - set
	os.Setenv("MY_VAR", "my value")
}