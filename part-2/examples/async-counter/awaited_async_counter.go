package main

import (
	"fmt"
	//"time"
    "sync"
)

func increment(counter *int, m *sync.Mutex, w *sync.WaitGroup) {
    m.Lock()
	for i := 0; i < 100; i++ {
		*counter = *counter + 1
	}
    m.Unlock()
    w.Done()
}

func decrement(counter *int, m *sync.Mutex, w *sync.WaitGroup) {
    m.Lock()
	for i := 0; i < 100; i++ {
		*counter = *counter - 1
	}
    m.Unlock()
    w.Done()
}

func main() {

	counter := 0
    var m sync.Mutex
    var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
        w.Add(1)
		go increment(&counter, &m, &w)
	}

	for i := 0; i < 1000; i++ {
        w.Add(1)
		go decrement(&counter, &m, &w)
	}

    // remove this nasty nasty line
    // time.Sleep(time.Second * 4)
    w.Wait()

    fmt.Printf("Counter = %v\n", counter)
}
