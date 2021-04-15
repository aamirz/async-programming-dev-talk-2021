package main

import (
	"fmt"
	"time"
    "sync"
)

func increment(counter *int, m *sync.Mutex) {
    m.Lock()
	for i := 0; i < 100; i++ {
		*counter = *counter + 1
	}
    m.Unlock()
}

func decrement(counter *int, m *sync.Mutex) {
    m.Lock()
	for i := 0; i < 100; i++ {
		*counter = *counter - 1
	}
    m.Unlock()
}

func main() {

	counter := 0
    var m sync.Mutex

	for i := 0; i < 1000; i++ {
		go increment(&counter, &m)
	}

	for i := 0; i < 1000; i++ {
		go decrement(&counter, &m)
	}

    time.Sleep(time.Second * 4)
	fmt.Printf("Counter = %v\n", counter)
}
