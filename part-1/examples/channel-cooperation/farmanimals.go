package main

import (
       "fmt"
       "sync"
)

func moo(messages chan string, w *sync.WaitGroup) {
     for i := 0; i < 10; i++ {
     messages <- "moo"
     }

    w.Done()
}

func quack(messages chan string, w *sync.WaitGroup) {
     for i := 0; i < 10; i++ {
     messages <- "quack"
     }

    w.Done()
}

func scree(messages chan string, w *sync.WaitGroup) {
     for i := 0; i < 10; i++ {
     messages <- "scree"
     }

    w.Done()
}


func megaphone(messages chan string, done chan bool) {
     fmt.Printf("playing voicemails from farmanimals to McMaster-Carr!\n")
     for {
     message, more  := <-messages
     if (more) {
        fmt.Printf("%s\n", message)
     } else {
        fmt.Println("printed all messages")
        break
}
}
    done <- true

}

func main() {
     var w sync.WaitGroup
     messages := make(chan string, 2)
     done := make(chan bool, 1)

     w.Add(3)
     go moo(messages, &w)
     go quack(messages, &w)
     go scree(messages, &w)

     go megaphone(messages, done)
     w.Wait()
     close(messages)

     <-done
}
