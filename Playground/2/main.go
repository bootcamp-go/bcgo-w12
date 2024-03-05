package main

// Anonymous function
func Anonymous() func() {
	return func() {
		println("Anonymous function")
	}
}

// Closure is a function that returns a function
// which holds a reference to a variable from the
// outer scope.
func Closure() func() {
	var i int
	return func() {
		i++
		println(i)
	}
}

func main() {
	// Anonymous function
	fnAn := Anonymous()
	fnAn()

	// Closure
	closure := Closure()
	closure()
	closure()
	closure()
}