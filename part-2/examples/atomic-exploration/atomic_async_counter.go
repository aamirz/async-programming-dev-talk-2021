package main

import (
	"fmt"
	//"time"
    "sync"
	"sync/atomic"
)


type Context struct {
	counter *int
	atomicFlag *int32
	w *sync.WaitGroup
}

func newContext(counter *int, flag *int32, w *sync.WaitGroup) *Context {
	context := Context{counter: counter, atomicFlag: flag, w: w}
	return &context
}

func increment(threadId int, context *Context) {

	// naive lock
	for ;; {
		if (atomic.CompareAndSwapInt32(context.atomicFlag, 0, 1)) {
			break;
		}
	}

	fmt.Printf("Thread %v incrementing\n", threadId)
	for i := 0; i < 100; i++ {
		*context.counter = *context.counter + 1
	}

	atomic.CompareAndSwapInt32(context.atomicFlag, 1, 0)

    context.w.Done()
}

func decrement(threadId int, context *Context) {

 	// naive lock
	for ;; {
		if (atomic.CompareAndSwapInt32(context.atomicFlag, 0, 1)) {
			break;
		}
	}

	fmt.Printf("Thread %v decrementing\n", threadId)
	for i := 0; i < 100; i++ {
		*context.counter = *context.counter - 1
	}

    atomic.CompareAndSwapInt32(context.atomicFlag, 1, 0)
    context.w.Done()
}

func main() {
	threads := 20
	counter := 0
    var w sync.WaitGroup
	var flag int32

	context := newContext(&counter, &flag, &w)

	threadId := 0
	for i := 0; i < threads/2; i++ {
        w.Add(1)
		go increment(threadId, context)
		threadId++
	}

	for i := 0; i < threads/2; i++ {
        w.Add(1)
		go decrement(threadId, context)
		threadId++
	}

    // remove this nasty nasty line
    // time.Sleep(time.Second * 4)
    w.Wait()

    fmt.Printf("Counter = %v\n", counter)
}
