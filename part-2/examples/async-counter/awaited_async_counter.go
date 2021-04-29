package main

import (
	"fmt"
	//"time"
    "sync"
)

func increment(threadId int, counter *int, m *sync.Mutex, w *sync.WaitGroup) {
    m.Lock()
	fmt.Printf("Thread %v incrementing\n", threadId)
	for i := 0; i < 100; i++ {
		*counter = *counter + 1
	}
    m.Unlock()
    w.Done()
}

func decrement(threadId int, counter *int, m *sync.Mutex, w *sync.WaitGroup) {
    m.Lock()
	fmt.Printf("Thread %v decrementing\n", threadId)
	for i := 0; i < 100; i++ {
		*counter = *counter - 1
	}
    m.Unlock()
    w.Done()
}

func main() {

	threads := 10
	counter := 0
    var m sync.Mutex
    var w sync.WaitGroup

	threadId := 0
	for i := 0; i < threads; i++ {
        w.Add(1)
		go increment(threadId, &counter, &m, &w)
		threadId++
	}

	for i := 0; i < threads; i++ {
        w.Add(1)
		go decrement(threadId, &counter, &m, &w)
		threadId++
	}

    // remove this nasty nasty line
    // time.Sleep(time.Second * 4)
    w.Wait()

    fmt.Printf("Counter = %v\n", counter)
}
