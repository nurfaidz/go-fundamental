package main

import "fmt"

func boolean() {
	// Declare and initialize a boolean variable
	var isRaining bool = true
	var isWeekend bool = false

	// Display initial values
	fmt.Println("Is it raining?", isRaining)
	fmt.Println("Is it weekend?", isWeekend)

	// Perform logical operations
	canGoOutside := !isRaining && isWeekend // Can go outside if it's not raining and it's weekend
	fmt.Println("Can I go outside?", canGoOutside)

	// Change the state of isRaining
	isRaining = false
	fmt.Println("Is it raining now?", isRaining)
}

func main() {
	boolean()
}
