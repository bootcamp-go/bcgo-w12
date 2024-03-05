package main

import (
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout

	// write
	bytes := []byte("Hello, Writer! How are you?")
	n, err := w.Write(bytes)
	if err != nil {
		println(err.Error())
		return
	}

	// print the number of bytes written
	println("\n", n, "bytes written")
}