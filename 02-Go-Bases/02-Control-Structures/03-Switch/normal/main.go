package main

func main() {
	// unique path condition
	var age int = 25

	switch {
	case age > 21:
		println("You can drink")
	case age > 18:
		println("You can vote")
	case age > 16:
		println("You can drive a car")
	default:
		println("not sure what you can do with that age")
	}

	// unique path condition
	var command string = "greetings"

	switch command {
	case "start":
		println("Starting the application")
	case "greetings", "salut":
		println("Hello, how are you?")
	case "stop":
		println("Stopping the application")
	default:
		println("Unknown command")
	}
}