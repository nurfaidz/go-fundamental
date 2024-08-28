package main

import (
	"fmt"
	"os"
)

func DeferAndExit() {
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Hei everyone")
	fmt.Println("Welcome back to Go Learning Center")
	fmt.Println()

	callDeferFunc()
	fmt.Println("Hei everyone")
	fmt.Println()

	// exit
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before exiting")
	os.Exit(1)
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("defer function starts to execute")
}

func main() {
	DeferAndExit()
}
