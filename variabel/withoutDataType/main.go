package main

import "fmt"

func implisit1() {
	var name = "JNE"
	var year = "2025"

	fmt.Println(name, year)
}

func implisit2() {
	name, year := "JNE", 2025

	fmt.Println(name, year)
}

func main() {
	implisit1()
	implisit2()
}
