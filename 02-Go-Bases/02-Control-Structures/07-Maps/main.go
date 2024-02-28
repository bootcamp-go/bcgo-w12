package main

import "fmt"

func main() {
	// map: initialize by default (zero value -> nil)
	// var fruitsCounter map[string]int

	// map: initialize native
	// fruitsCounter := map[string]int{}

	// map: initialize native
	// fruitsCounter := map[string]int{
	// 	"apple":  2,
	// 	"banana": 3,
	// 	"orange": 4,
	// }

	// map: make
	fruitsCounter := make(map[string]int)
	fruitsCounter["apple"] = 2
	fruitsCounter["banana"] = 3
	fruitsCounter["orange"] = 4
	fmt.Println("orange:", fruitsCounter["orange"])

	// delete key
	delete(fruitsCounter, "banana")

	fmt.Println("fruitsCounter:", fruitsCounter)

	// check key exists
	mango, ok := fruitsCounter["mango"]
	if ok {
		fmt.Println("mango:", mango)
	} else {
		fmt.Println("mango not exists")
	}

	// mango = fruitsCounter["mango"]
	// if mango > 0 {
	// 	fmt.Println("mango:", mango)
	// } else {
	// 	fmt.Println("mango not exists")
	// }
}