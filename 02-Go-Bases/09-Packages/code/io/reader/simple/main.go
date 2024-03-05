package main

import (
	"io"
	"strings"
)

func main() {
	var r io.Reader
	r = strings.NewReader("Hello, Reader! How are you?")

	// create a buffer to read into
	var buf = make([]byte, 64)

	// read into the buffer
	// - one read
	n, err := r.Read(buf)
	if err != nil {
		println(err.Error())
		return
	}

	// print the number of bytes read and the buffer
	println(n, string(buf[:n]))

	// try to read again
	// - no more data to read
	buf = make([]byte, 64)
	n, err = r.Read(buf)
	if err != nil {
		// expect an EOF error
		println(err.Error())
		return
	}

	// print the number of bytes read and the buffer
	println(n, string(buf[:n]))
}