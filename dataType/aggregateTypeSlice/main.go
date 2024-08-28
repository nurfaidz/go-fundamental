package main

import "fmt"

func main() {
	// Declarea and initialize an array
	arr := [6]int{1, 2, 3, 4, 5, 6}

	// Make a slice from the array
	s1 := arr[1:4] // Take elements from index 1 to 3
	fmt.Println("Slice 1:", s1)
	fmt.Println("Length of slice 1: %d\n", len(s1))
	fmt.Println("Capacity of slice 1: %d\n", cap(s1))

	// Change the value of the slice
	s1[0] = 20
	fmt.Println("Array after modifying slice 1:", arr)
	fmt.Println("Updated slice 1:", s1)

	// Make slice with make function
	s2 := make([]int, 3, 5) // Make a slice with length 3 and capacity 5
	fmt.Println("\nSlice 2:", s2)
	fmt.Println("Length of slice 2: %d\n", len(s2))
	fmt.Println("Capacity of slice 2: %d\n", cap(s2))

	// Set values to the slice
	s2[0] = 10
	s2[1] = 20
	s2[2] = 30
	fmt.Println("Updated slice 2:", s2)

	// Add elements to the slice using append
	s2 = append(s2, 40, 50)
	fmt.Println("\nSlice 2 after append:", s2)
	fmt.Println("Length of slice 2 after append: %d\n", len(s2))
	fmt.Println("Capacity of slice 2 after append: %d\n", cap(s2))
}
