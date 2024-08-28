package main

import (
	"fmt"
	"time"
)

func main() {
	// run  func`pintMessage` in goroutine
	go printMessage("Hello from Goroutine")

	fmt.Println("Hello from main")

	// waiting some seconds to finish goroutine
	time.Sleep(2 * time.Second)
}

// func of running in goroutine
func printMessage(message string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println(message)
}
