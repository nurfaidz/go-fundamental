package main

import "fmt"

func main() {
	x := 1
	if x < 0 {
		panic("Nilai x harus positif")
	} else {
		fmt.Println("Nilai x adalah", x)
	}
}
