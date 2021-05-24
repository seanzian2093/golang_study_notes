package main

import (
	"fmt"
	"sync"
	"time"
)

// a WaitGroup must be passed to function by pointer
func worker(id int, wg *sync.WaitGroup) {
	// notify WaitGroup that we are done
	defer wg.Done()
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block untill all workers have notified WaitGroup that we are done.
	wg.Wait()
}