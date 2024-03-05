package main

import (
	"io"
	"strings"
)

func main() {
	// reader
	var r io.ReadSeeker
	r = strings.NewReader("Hello, Reader! How are you?")

	// read into the buffer
	for {
		// create a buffer
		buf := make([]byte, 8)

		// read
		n, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			println(err.Error())
			return
		}

		// print the buffer
		println(string(buf[:n]))
	}

	// seek
	_, err := r.Seek(0, io.SeekStart)
	if err != nil {
		println(err.Error())
		return
	}

	// one read
	buf := make([]byte, 8)
	n, err := r.Read(buf)
	if err != nil {
		println(err.Error())
		return
	}

	// print the buffer
	println(string(buf[:n]))
}