// Package ppsync provides basic mutual exclusion lock primitives.
package ppsync

import (
	"sync"
	"sync/atomic"
)

// A ALock is a mutual exclusion lock that represents an anderson lock.
// The zero value for a ALock is an unlocked mutex.
type ALock struct {
	queue        []bool
	acquireIndex *int64 // the next index in the queue for a locking thread
	nextIndex    *int64 // the index of the next thread to be handed the lock
}

// NewALock creates, initializes and returns a new ALock object.
// nThreads is the number of threads in the thread pool -- if more
// threads than this number attempt to access the lock, then there is
// a runtime error in the form of a panic().
func NewALock(nThreads int64) sync.Locker {
	queue := make([]bool, nThreads)
	var acquireIndex int64
	var nextIndex int64
	// allow first thread to proceed
	queue[0] = true
	acquireIndex = -1
	nextIndex = 1
	return &ALock{queue, &acquireIndex, &nextIndex}
}

// Lock locks lock. If the lock is already in use, the calling goroutine
// blocks until the lock is available.
func (lock *ALock) Lock() {
	n := int64(len(lock.queue))
	// Thread local index for the current thread.
	myIndex := atomic.AddInt64(lock.acquireIndex, 1)

	// wait on the lock to become available
	for !lock.queue[myIndex%n] {
	}

	// reset my own position, set position to unlock
	*lock.nextIndex = (myIndex + 1) % n
	lock.queue[myIndex%n] = false
}

// Unlock unlocks lock.
// It is a run-time error if lock is not locked on entry to Unlock.
//
// A locked lock is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a lock and then
// arrange for another goroutine to unlock it.
func (lock *ALock) Unlock() {
	lock.queue[*lock.nextIndex] = true
}
