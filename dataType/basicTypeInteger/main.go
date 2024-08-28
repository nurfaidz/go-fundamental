package main

import "fmt"

func main() {
	// Declare and initialize integers of different types
	var a int = 42                    // Default integer (usually 64 bits on a 64-bit systems)
	var b int8 = 12                   // 8-bit signed integer
	var c int16 = -1000               // 16-bit signed integer
	var d int32 = 100000              // 32-bit signed integer
	var e int64 = 1234567890123456789 // 64-bit signed integer

	// Perform some basic operations
	sum := a + int(b) + int(c) + int(d) + int(e) // Smm of all integers will cause mismatch error and convert d to int64 for subtraction

	// difference := e - d
	different := e - int64(d)

	fmt.Println("Sum:", sum)

	fmt.Println("Different:", different)
}
