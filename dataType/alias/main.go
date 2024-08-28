package main

import "fmt"

func main() {
	// Declare variables of different basic data types
	var a byte = 255 // byte(uint8)
	var b rune = 'G' // rune(int32), representing the Unicode code point for 'G'
	var c int = -42  // int(signed integer)
	var d uint = 42  // uint(unsigned integer)

	// Display the values and types of each variable
	fmt.Printf("Value of byte a: %d, Type: %T\n", a, a)
	fmt.Printf("Value of rune b: %c, Type: %T\n", b, b) // %c to print the character
	fmt.Printf("Value of int c: %d, Type: %T\n", c, c)
	fmt.Printf("Value of uint d: %d, Type: %T\n", d, d)
}
