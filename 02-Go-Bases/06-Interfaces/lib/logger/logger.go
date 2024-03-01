package logger

// Logger is a simple interface to log messages at different levels.
type Logger interface {
	// Info logs a message at level Info on the standard logger.
	Info(format string, args ...any)
}