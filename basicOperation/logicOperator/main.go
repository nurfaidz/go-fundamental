package main

import "fmt"

func LogicOperator() {
	a := true
	b := false

	andResult := a && b
	fmt.Println("a && b : %v\n", andResult)

	orResult := a || b
	fmt.Println("a || b : %v\n", orResult)

	notA := !a
	notB := !b
	fmt.Println("!a: %v, !b: %v", notA, notB)

	// Example using of conditional operator
	if a && !b {
		fmt.Println("a adalah true dan b adalah false")
	}

	if a || b {
		fmt.Println("Setidaknya satu dari a atau b adalah true")
	}
}

func main() {
	LogicOperator()
}
