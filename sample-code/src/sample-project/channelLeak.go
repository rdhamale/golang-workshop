package main

import (
	"fmt"
	"time"
)

var n int = 100

// Task ...
func Task(i int) {
	fmt.Println("Box", i)
}

// DemonstrateLeak ...
func DemonstrateLeak() {
	ack := make(chan int, n) // Acknowledgement channel

	for i := 0; i < 50; i++ {
		go func() {
			Task(<-ack)
		}()
	}

	for i := 0; i < 1000; i++ {
		go func(i int) {
			ack <- i
		}(i)
	}
	time.Sleep(60 * time.Second)
}
