package logger

// NewDummy creates a new dummy logger
func NewDummy() *Dummy {
	return &Dummy{}
}

// Dummy is a dummy logger that does nothing
type Dummy struct{}

func (d Dummy) Logf(format string, args ...interface{}) {
	// Do nothing
}

func (d Dummy) Warnf(format string, args ...interface{}) {
	// Do nothing
}