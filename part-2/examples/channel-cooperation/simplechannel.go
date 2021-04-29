package main

import (
    "fmt"
    "bufio"
    "os"
    "sync"
)

func printMessage(messages chan string, w *sync.WaitGroup) {
    message := <-messages
    fmt.Printf("The message is: %s", message)
    w.Done()
}

// will this code work?
func main() {
    var w sync.WaitGroup
    messages := make(chan string, 1)

    // wait until a message is available then print it
    w.Add(1)
    go printMessage(messages, &w)

    // accept a line of input
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    reader.ReadString('\n')

    // await the message write before termination
    w.Wait()
}
