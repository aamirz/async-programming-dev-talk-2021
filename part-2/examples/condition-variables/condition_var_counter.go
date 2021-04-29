package main

import (
	"fmt"
	//"time"
    "sync"
)

func increment(threadId int, counter *int, m *sync.Mutex, cond *sync.Cond, w *sync.WaitGroup) {
    m.Lock()
	fmt.Printf("Thread %v incrementing\n", threadId)
	for i := 0; i < 100; i++ {
		*counter = *counter + 1
	}

	if (*counter > 500) {
		cond.Broadcast()
	}

	// wait until the counter has achieved 4
	for *counter < 500 {
		cond.Wait()
	}
	fmt.Printf("Thread %v has finished waiting\n", threadId)
	m.Unlock()

	// condition variables
    w.Done()
}

func main() {

	threads := 10
	counter := 0
    var m sync.Mutex
    var w sync.WaitGroup

	condVar := sync.NewCond(&m)

	threadId := 0
	for i := 0; i < threads; i++ {
        w.Add(1)
		go increment(threadId, &counter, &m, condVar, &w)
		threadId++
	}


    // remove this nasty nasty line
    // time.Sleep(time.Second * 4)
    w.Wait()

    fmt.Printf("Counter = %v\n", counter)
}
