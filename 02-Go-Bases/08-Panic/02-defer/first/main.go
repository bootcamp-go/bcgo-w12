package main

import (
	"io"
	"os"
)

func main() {
	src := "./02-defer/first/data.txt"
	dst := "./02-defer/first/data_copy.txt"

	defer func() {
		println("I'm the last defer.")
	}()

	TransferData(src, dst)
}

// TransferData transfers data from one file to another.
func TransferData(src, dst string) {
	// Open the source file.
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	// defer srcFile.Close()
	defer func (f *os.File) {
		println("Closing the source file.")
		srcFile.Close()
	}(srcFile)

	// Open the destination file.
	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	// defer dstFile.Close()
	defer func () {
		println("Closing the destination file.")
		dstFile.Close()
	}()

	panic("Something went wrong.")

	// Transfer the data.
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		panic(err)
	}

	defer func () {
		println("I'm the last defer.")
	}()

	println("Data transferred successfully.")

	// BAD PRACTICE: Don't left the resource closing at the end of the function.
	// use defer instead right after the resource is opened.
	// ---
	// // Close the source file.
	// srcFile.Close()
	// // Close the destination file.
	// dstFile.Close()
}