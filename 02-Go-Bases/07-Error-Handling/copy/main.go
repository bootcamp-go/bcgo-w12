package main

var (
	// ErrRandom is the error message when something random happens
	ErrRandom = New("something random happened")
)

// New returns a new error that formats as the given text.
func New(text string) error {
	return CustomError{text}
}

type CustomError struct {
	msg string
}

func (e CustomError) Error() string {
	return e.msg
}

func main() {
	if ErrRandom == New("something random happened") {
		println("Error messages are equal")
		return
	}

	println("Error messages are different")
}