package main

import "fmt"

func main() {
	// String menggunakan tanda petik dua
	str1 := "Ini adalah string menggunakan petik dua."
	fmt.Println("String 1:", str1)

	// String dengan karakter newline dan tab
	str2 := "Ini adalah string\nmengandung newline\n\tdan tab."
	fmt.Println("String 2:", str2)

	// String dengan karakter backtick
	str3 := `Ini adalah string menggunakan backtick. Karakter newline dan tab tidak di-escape di sini:\tTab\nNewline`
	fmt.Println("String 3:", str3)

}
