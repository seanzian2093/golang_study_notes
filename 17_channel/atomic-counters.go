package main

import (
	"fmt"
	"sync"
	// "sync/atomic"
)

func main() {
	// define ops as unsigned integer, i.e. always positive
	var ops uint64
	var wg sync.WaitGroup

	for i := 1; i <= 50; i++ {
		wg.Add(1)

		go func() {
			for j := 1; j <= 1000; j++ {
				// with atomic.AddUint64() always 5000
				// atomic.AddUint64(&ops, 1)

				// with regular incrment, result is different between each run since goroutines interfere with each other
				ops++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops: ", ops)
}