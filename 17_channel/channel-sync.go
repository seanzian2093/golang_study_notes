package main

import (
    "fmt"
    "time"
)

// worker will be run in a goroutine
// done is a channel that is used to notify other goroutine that this worker is done.
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

	// send a value to channel done that it is done.
    done <- true
}

func main() {

	// build a channel and provide to worker
    done := make(chan bool, 1)
    go worker(done)

	// Block untill we receive a notification from worker through channel done.
	// If this statement is removed; 
	// main() will proceed without wait for worker or  before worker even started?
    <-done
}