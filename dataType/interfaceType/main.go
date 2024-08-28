package main

import "fmt"

// Define an interface using Speak method
type Speaker interface {
	Speak() string
}

// Type implementation of Speaker interface
type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

type Dog struct {
	Breed string
}

func (d Dog) Speak() string {
	return "Woof! I am a " + d.Breed
}

// Function that takes a Speaker interface
func greet(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	alice := Person{Name: "Alice"}
	buddy := Dog{Breed: "Golden Retriever"}

	greet(alice)
	greet(buddy)

	// The nil interface value
	var i interface{} = 42
	fmt.Println("Value of i:", i)

	// Type assertion
	if v, ok := i.(int); ok {
		fmt.Println("i is of type int with value:", v)
	} else {
		fmt.Println("i is not of type int")
	}
}
