package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Run()
}

func Run() {
	// Process the files.
	files := []string{
		"./03-recover/first/data_1.txt",
		"./03-recover/first/data_2.txt",
		"./03-recover/first/data_3.txt",
	}
	
	for ix, file := range files {
		err := processFile(file)
		if err != nil {
			fmt.Printf("File %d - Path %s - Error %v\n", ix, file, err)
			continue
		}

		fmt.Printf("File %d - Path %s processed.\n", ix, file)
	}
}

func processFile(path string) (err error) {
	// recover
	defer func () {
		r := recover(); if r != nil {
			err = fmt.Errorf("caught a panic. Error: %v", r)
			return
		}
	}()

	// Open the file.
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file.
	_, err = io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// Process the file.
	// ...
	return
}