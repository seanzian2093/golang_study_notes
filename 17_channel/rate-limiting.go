package main

import (
	"fmt"
	"time"
)


func main() {
	// suppose we want to limit our handling of incoming requests we could serve them off a channel
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		// receiving from a channel is blocking, so we handle request every 200 m sec
		<- limiter
		fmt.Println("request ", req, time.Now())
	}

	// suppose we allow a short burst of 3 incoming requets being handled at one time

	// Initialize a buffered channel to store busrty requests
	// use time.Time to track handling time stamp
	burstylimiter := make(chan time.Time, 3)
	for i := 1; i <= 3; i++ {
		burstylimiter <- time.Now()
	}
	// then try to fill the burstylimiter, up to 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstylimiter <- t
		}
	}()

	// suppose now we have 5 incoming requets when bursting is allowed
	burstyrequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyrequests <- i
	}
	close(burstyrequests)

	for req := range burstyrequests {
		// up to 3 values can be received from burstylimiter at one time.
		<- burstylimiter
		fmt.Println("request ", req, time.Now())
	}
}