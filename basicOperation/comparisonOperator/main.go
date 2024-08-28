package main

import "fmt"

func ComparisonOperator() {
	a := 10
	b := 20
	c := 30

	fmt.Println(" a > b:", a > b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= c:", a <= c)
	fmt.Println("a == c:", a == c)
	fmt.Println("a != b:", a != b)
}

func main() {
	ComparisonOperator()
}
