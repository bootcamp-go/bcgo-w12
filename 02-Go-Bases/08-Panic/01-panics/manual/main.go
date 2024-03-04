package main

import (
	"io"
	"os"
)

func main() {
	data := ReadFileData("data.txt")
	println(data)
}

// ReadFileData reads the content of a file and returns it as a string.
func ReadFileData(path string) (data string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	data = string(bytes)
	return
}