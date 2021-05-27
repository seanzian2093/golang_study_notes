package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {

	// Go signal notification works by sending os.Signal values on a channel. 
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

	// signal.Notify registers the given channel to receive notifications of the specified signals
	// Notify accepts variadic arguments of syscall signals, here are two of them
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}