package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println("The sunday is", Sunday)
	fmt.Println("The monday is", Monday)
	fmt.Println("The tuesday is", Tuesday)
	fmt.Println("The wednesday is", Wednesday)
	fmt.Println("The thursday is", Thursday)
	fmt.Println("The friday is", Friday)
	fmt.Println("The saturday is", Saturday)
}
