package main

import (
	"fmt"
)

// DemonstrateChannels ...
func DemonstrateChannels() {
	out := make(chan int)
	in := make(chan int)

	for i := 0; i < 100; i++ {
		go multiplyByTwo(in, out)
		in <- i
		fmt.Println(<-out)

	}
}

func multiplyByTwo(in <-chan int, out chan<- int) {
	fmt.Println("Initializing goroutine...")
	num := <-in
	result := num * 2
	out <- result
}
