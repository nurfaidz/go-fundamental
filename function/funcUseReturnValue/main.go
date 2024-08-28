package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func FuncUseReturnValue() {
	result := getHello("Alice")
	fmt.Println(result)
}

func main() {
	FuncUseReturnValue()
}
