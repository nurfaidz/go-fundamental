package main

import "fmt"

func main() {
	// Declare and initialize map using string as key and int as value
	m := make(map[string]int)

	// Add key-value pairs to the map
	m["apple"] = 10
	m["banana"] = 20
	m["cherry"] = 30

	// Print the map
	fmt.Println(m)

	// Access the value of a key
	fmt.Println("Values of banana:", m["banana"])

	// Delete element from the map
	delete(m, "banana")
	fmt.Println("Map after deleting 'banana':", m)

	// Check if a key exists in the map
	value, exists := m["cherry"]
	if exists {
		fmt.Println("Value from 'cherry':", value)
	} else {
		fmt.Println("'cherry' not found in map")
	}

	// Add element using of same key will update the before value
	m["apple"] = 100
	fmt.Println("Map after updating 'apple':", m)
}
