package main

import (
	"fmt"
	"time"
	"math/rand"
)

/* */
/* Insead of using a WaitGroup we will use */
/* channels to signal termination.
 */

func worker(done chan bool, id int, howLong time.Duration) {
	fmt.Printf("Thread %v starting...\n", id)
	time.Sleep(howLong)
	fmt.Printf("Thread %v finished.\n", id)
}

func main() {
	done := make(chan bool, 2)

	n := 10
	rand.Seed(time.Now().UnixNano())

	waitThreadA := time.Duration(rand.Intn(n)) * time.Second
	waitThreadB := time.Duration(rand.Intn(n)) * time.Second

	go worker(done, 1, waitThreadA)
	go worker(done, 2, waitThreadB)

	<-done
	<-done
}
