package main

import (
	"fmt"
	"time"
)

func main() {
	// create a some channel
	ch1 := make(chan string)
	ch2 := make(chan string)

	// func goroutine for ch1
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Data from ch1"
	}()

	// func goroutine for ch2
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Data from ch2"
	}()

	// use select to received data from the channel that is ready
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}
