package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func Function() {
	sayHello("Alice")
}

func main() {
	Function()
}
