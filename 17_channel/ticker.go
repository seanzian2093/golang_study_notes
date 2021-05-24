package main

import (
	"fmt"
	"time"
)

func main() {
	// ticker for doing sth repeatedly at regular interval in future;
	// timer for once in future
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// an infinite loop in a goroutine
	// loop only breaks when a value is sent to done
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick: ",t)
			}
		}
	}()

	// synchronically blocking 1600 m sec
	time.Sleep(1600 * time.Millisecond)
	// stop the ticker
	ticker.Stop()
	// send a value to done
	done <-true
	// even with more time, ticker will not fire since it is stopped
	time.Sleep(1600 * time.Millisecond)
	fmt.Println("Ticker stopped")
}