package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	// Read data from file
	file, err := os.OpenFile("./code/os/file-read/test.txt", os.O_RDONLY, 0)
	if err != nil {
		println(err.Error())
		return
	}
	defer file.Close()

	// read by line
	rd := bufio.NewReader(file)
	for {
		data, _, err := rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			println(err.Error())
			return
		}

		// apply necessary operation to data
		println(string(data))
	}
}