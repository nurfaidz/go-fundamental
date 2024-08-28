package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func FuncWithParam() {
	sayHelloTo("Alice", "Bob")
}

func main() {
	FuncWithParam()
}
