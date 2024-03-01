package main

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	// Name string  // this case is not ambiguous
	A
	B
}

func main() {
	c := C{
		A: A{
			Name: "A",
		},
		B: B{
			Name: "B",
		},
	}
	println(c.A.Name)
	// println(c.Name) // ambiguous selector c.Name
}