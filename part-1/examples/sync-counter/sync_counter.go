package main

import (
	"fmt"
	"time"
)

func increment(counter *int) {
	for i := 0; i < 1000000; i++ {
		*counter = *counter + 1
	}
}

func decrement(counter *int) {

	for i := 0; i < 1000000; i++ {
		*counter = *counter - 1
	}
}

func main() {

	counter := 0

	for i := 0; i < 1000; i++ {
		increment(&counter)
	}

	for i := 0; i < 1000; i++ {
		decrement(&counter)
	}

	time.Sleep(time.Second * 5)
	fmt.Printf("Counter = %v\n", counter)
}
