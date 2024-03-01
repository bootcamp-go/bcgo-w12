package logger

import (
	"fmt"
	"os"
)

// NewTextFile creates a new TextFile logger with the given path.
func NewTextFile(path string) TextFile {
	return TextFile{Path: path}
}

// TextFile is a simple logger that logs messages at different levels to a file.
type TextFile struct {
	// Path is the path to the file where the logs will be written.
	Path string
}

// Info logs a message at level Info to the file.
func (l TextFile) Info(format string, args ...any) {
	// open file
	file, err := os.OpenFile(l.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// write to file
	file.WriteString(fmt.Sprintf(format, args...) + "\n")
}