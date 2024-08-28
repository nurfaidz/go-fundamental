package main

import "fmt"

// Declare a struct type with name `Person`
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Make instance of `Person` struct
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}

	// Display the values fields of the 'Person' struct
	fmt.Println("Person 1:")
	fmt.Println("First Name:", p1.FirstName)
	fmt.Println("Last Name:", p1.LastName)
	fmt.Println("Age:", p1.Age)

	// Access and change the field of the struct
	p1.Age = 31
	fmt.Println("Updated Age:", p1.Age)
}
