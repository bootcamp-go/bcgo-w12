package main

func main() {
	var cc int
	var oddNumbers int

	for cc < 100 {
		// increment counter
		cc += 3

		// check if even number
		if cc%2 == 0 {
			continue
		}

		// increment odd numbers
		oddNumbers++
	}

	println("Odd numbers:", oddNumbers)
}