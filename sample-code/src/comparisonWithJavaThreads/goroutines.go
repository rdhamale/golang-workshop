package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println("Total time: ", time.Since(now))
	}()
	out := make(chan int)
	for i := 0; i < 1000000; i++ {
		go print(i, out)
		fmt.Print(<-out)
	}

}

func print(i int, out chan<- int) {
	out <- i
	fmt.Print(" goroutine: is running\n")
}
