package main

import "fmt"

func Arithmetic() {
	a := 10
	b := 5
	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	modulus := a % b

	fmt.Println("%d + %d = %d\n", a, b, tambah)
	fmt.Println("%d - %d = %d\n", a, b, kurang)
	fmt.Println("%d * %d = %d\n", a, b, kali)
	fmt.Println("%d / %d = %d\n", a, b, bagi)
	fmt.Println("%d %% %d = %d\n", a, b, modulus)

	a++
	fmt.Println("Increment: %d\n", a) // Increment
	b--
	fmt.Println("Decrement: %d\n", b) // Decrement
}

func main() {
	Arithmetic()
}
