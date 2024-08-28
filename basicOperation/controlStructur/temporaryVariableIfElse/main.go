package main

import "fmt"

func main() {
	kodeKota := "CGK"

	if length := len(kodeKota); length > 3 {
		fmt.Println("Kode kota terlalu panjang")
	} else {
		fmt.Println("Kode kota sudah benar!")
	}
}
