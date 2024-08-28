package main

import "fmt"

// Define a function with a name "greet" no return value
func greet(name string) {
	fmt.Println("Hello,", name)
}

// Define a function with a name "add" return an two integer value
func add(a int, b int) int {
	return a + b
}

func Function() {
	// Call the function greet
	greet("Adam")

	// Call the function add and save the return value
	sum := add(5, 3)
	fmt.Println("Sum:", sum)
}

func main() {
	Function()
}
