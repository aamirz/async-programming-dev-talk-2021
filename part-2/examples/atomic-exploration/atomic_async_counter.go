package main

import (
	"fmt"
	//"time"
    "sync"
	"sync/atomic"
)


func increment(threadId int,
	counter *int, atomicFlag *int32, w *sync.WaitGroup) {

	// naive lock
	for ;; {
		if (atomic.CompareAndSwapInt32(atomicFlag, 0, 1)) {
			break;
		}
	}

	fmt.Printf("Thread %v incrementing\n", threadId)
	for i := 0; i < 100; i++ {
		*counter = *counter + 1
	}

	atomic.CompareAndSwapInt32(atomicFlag, 1, 0)

    w.Done()
}

func decrement(threadId int,
	counter *int, atomicFlag *int32,  w *sync.WaitGroup) {

 	// naive lock
	for ;; {
		if (atomic.CompareAndSwapInt32(atomicFlag, 0, 1)) {
			break;
		}
	}

	fmt.Printf("Thread %v decrementing\n", threadId)
	for i := 0; i < 100; i++ {
		*counter = *counter - 1
	}

    atomic.CompareAndSwapInt32(atomicFlag, 1, 0)
    w.Done()
}

func main() {

	threads := 10
	counter := 0
    var w sync.WaitGroup
	var atomicFlag int32
	threadId := 0
	for i := 0; i < threads; i++ {
        w.Add(1)
		go increment(threadId, &counter, &atomicFlag, &w)
		threadId++
	}

	for i := 0; i < threads; i++ {
        w.Add(1)
		go decrement(threadId, &counter, &atomicFlag, &w)
		threadId++
	}

    // remove this nasty nasty line
    // time.Sleep(time.Second * 4)
    w.Wait()

    fmt.Printf("Counter = %v\n", counter)
}
