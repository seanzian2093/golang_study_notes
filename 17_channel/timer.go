package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	// timer1.C is a channel that will be notified when a timer fires.
	// timer1.C blocks untill it sends a value which notifying the timer fired.
	<- timer1.C
	// so this will be printed to console after about 2 seconds
	fmt.Println("timer1 fired.")

	timer2 := time.NewTimer(1 * time.Second)
	// put a timer in a goroutine
	go func() {
		<- timer2.C
		fmt.Println("timer2 fired.")
	}()

	// synchronically stop it
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stopped.")
	}

	// wait enough time so the timer2 could have fired
	// but will not since it is stopped before fire
	time.Sleep(2 * time.Second)
}