package logger

import (
	"github.com/fatih/color"
)

// NewLocal creates a new Local logger with the given color.
func NewLocal(color string) Local {
	return Local{color: color}
}

// Local is a simple logger that logs messages at different levels.
type Local struct {
	color string
}

func (l Local) Info(format string, args ...any) {
	switch l.color {
	case "red":
		color.Red(format, args...)
	case "green":
		color.Green(format, args...)
	default:
		color.White(format, args...)
	}
}