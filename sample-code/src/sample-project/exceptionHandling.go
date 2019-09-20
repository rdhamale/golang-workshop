package main

import (
	"fmt"
	"log"
	"strconv"
)

// DemoException ...
func DemoException() {

	stringVal := "abc123"
	intVal, err := strconv.ParseInt(stringVal, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("integer Val: %d", intVal)
}
