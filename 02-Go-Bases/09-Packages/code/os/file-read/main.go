package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	// Read data from file
	file, err := os.OpenFile("./code/os/file-read/test.txt", os.O_RDONLY, 0644)
	if err != nil {
		println(err.Error())
		return
	}
	defer file.Close()

	for {
		// read data step by step
		buffer := make([]byte, 16)

		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			println(err.Error())
			return
		}

		// apply necessary operation to buffer[:n]
		println(string(buffer[:n]))
	}

	// reset file pointer to the beginning of file
	file.Seek(0, io.SeekStart)

	// read by line
	rd := bufio.NewReader(file)

	for {
		data, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			println(err.Error())
			return
		}

		// apply necessary operation to data
		println(data)
	}

	// read all data at once
	// bytes, err := io.ReadAll(file)
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }

	// println(string(bytes))
}