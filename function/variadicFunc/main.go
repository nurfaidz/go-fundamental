package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func VariadicFunc() {
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)
}

func main() {
	VariadicFunc()
}
