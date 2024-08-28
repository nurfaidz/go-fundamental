package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi panic:", r)
		}
	}()

	panic("Ups, ada kesalahan!")
	fmt.Println("Kode setelah panic tidak akan dijalankan")
}
