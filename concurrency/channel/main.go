package main

import (
	"fmt"
	"time"
)

func main() {
	// make channel using 2 capacity
	ch := make(chan string, 2)

	// func goroutin to send data to channel
	go func() {
		ch <- "Hello" // send data to channel
		ch <- "World"
		close(ch) // close channel
	}()

	// to receive data from channel
	for msg := range ch {
		fmt.Println("Received: ", msg)
	}

	// waiting some seconds goroutin to finish
	time.Sleep(2 * time.Second)
}
