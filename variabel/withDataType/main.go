package main

import "fmt"

func ReassignVariable() {
	var name string
	var year int = 2024

	name = "JNE"
	year = 2025

	fmt.Println(name, year)
}

func main() {
	ReassignVariable()
}
