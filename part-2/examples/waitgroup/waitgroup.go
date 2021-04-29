package waitgroup

import (
	"runtime"
	"sync/atomic"
)

// From the Go Documentation: A WaitGroup waits for a collection of
// goroutines to finish. The main goroutine calls Add to set the number
// of goroutines to wait for.  Then each of the goroutines runs and calls
// Done when finished.  At the same time, Wait can be used to block until
// all goroutines have finished.
type WaitGroup interface {
	Add(amount uint)
	Done()
	Wait()
}

/*
* The shared state that controls the internal synchronization of the
* WaitGroup.
 */
type sharedWaitState struct {
	threadNum          *uint64
	accessSynchronizer *int32
}

func (state *sharedWaitState) Add(amount uint) {
	for !atomic.CompareAndSwapInt32(state.accessSynchronizer, 0, 1) {
		runtime.Gosched()
	}

	*(state.threadNum) = *(state.threadNum) + uint64(amount)

	atomic.StoreInt32(state.accessSynchronizer, 0)
}

func (state *sharedWaitState) Done() {
	for !atomic.CompareAndSwapInt32(state.accessSynchronizer, 0, 1) {
		runtime.Gosched()
	}

	if atomic.LoadUint64(state.threadNum) > 0 {
		*(state.threadNum) = *(state.threadNum) - 1
	} else {
		panic("Attempted to call Done on a thread that has not been Added to the WaitGroup")
	}

	atomic.StoreInt32(state.accessSynchronizer, 0)
}

func (state *sharedWaitState) Wait() {
	for !(atomic.LoadUint64(state.threadNum) == 0) {
		runtime.Gosched()
	}
}

// NewWaitGroup returns a instance of a waitgroup
// This instance must be a pointer and should not
// be copied after creation.
func NewWaitGroup() WaitGroup {
	var threads uint64
	var access int32

	var wg = sharedWaitState{threadNum: &threads,
		accessSynchronizer: &access}

	return &wg
}
