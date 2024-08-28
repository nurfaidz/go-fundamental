package main

import "fmt"

func main() {
	// Declare integer variable
	num := 10
	fmt.Println("Original value of num: ", num)

	// Get the address of num
	ptr := &num
	fmt.Println("Pointer to num:", ptr)
	fmt.Println("Value at pointer address:", *ptr)

	// Change the value with pointer
	*ptr = 20
	fmt.Println("Updated value of num:", num)
	fmt.Println("Value at pointer address after update:", *ptr)

	// Using pointer in function
	increment(&num)
	fmt.Println("Value of num after increment function:", num)
}

// Function to increment the value of integer
func increment(n *int) {
	*n++ // Change value in address from pointer
}
