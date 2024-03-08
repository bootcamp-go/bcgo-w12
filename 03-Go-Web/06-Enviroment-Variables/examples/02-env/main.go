package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Enviroment variables
	// - load from file
	// loadenv(".env")
	err := godotenv.Load() // .env
	if err != nil {
		println("Error loading .env file")
		return
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	filePath := os.Getenv("FILE_PATH")

	println("HOST:", host)
	println("PORT:", port)
	println("FILE_PATH:", filePath)
}

// func loadenv(path string) {
// 	// open file
// 	f, err := os.Open(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	// bufio reader
// 	rd := bufio.NewReader(f)
// 	for {
// 		// read line
// 		line, _, err := rd.ReadLine()
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			panic(err)
// 		}
// 		// split line
// 		elements := strings.Split(string(line), "=")
// 		// set environment variable
// 		os.Setenv(elements[0], elements[1])
// 	}

// }