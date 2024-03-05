package main

import "os"

func main() {
	file, err := os.OpenFile("./code/os/file-write/test.txt", os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		println("open:", err.Error())
		return
	}
	defer file.Close()

	// Write data to file
	data := []byte("Hello, World!")

	_, err = file.Write(data)
	if err != nil {
		println("write:", err.Error())
		return
	}

	println("Write data to file success")
}