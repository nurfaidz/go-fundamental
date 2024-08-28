package main

import "fmt"

func main() {
	// Define and call the anonymous function
	func() {
		fmt.Println("This is an anonymous function")
	}() // Output: This is an anonymous function

	// Save the anonymous function to a variable and call it
	greet := func(name string) {
		fmt.Println("Hello " + name)
	}
	greet("Alice") // Output: Hello Alice

	// Use anonymous function as an argument to another function
	result := func(x int, y int) int {
		return x + y
	}(5, 3)
	fmt.Println("Multiplication result:", result) // Output: Multiplication result: 8
}
