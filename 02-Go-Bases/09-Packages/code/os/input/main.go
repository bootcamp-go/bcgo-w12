package main

import "os"

func main() {
	// read from stdin
	buffer := make([]byte, 16)

	n, err := os.Stdin.Read(buffer)
	if err != nil {
		println(err.Error())
		return
	}

	println(string(buffer[:n]))

	// read again
	n, err = os.Stdin.Read(buffer)
	if err != nil {
		println(err.Error())
		return
	}

	println(string(buffer[:n]))

	// for {
	// 	buffer := make([]byte, 16)

	// 	n, err := os.Stdin.Read(buffer)
	// 	if err != nil {
	// 		println(err.Error())
	// 		return
	// 	}

	// 	println(string(buffer[:n]))
	// }
}