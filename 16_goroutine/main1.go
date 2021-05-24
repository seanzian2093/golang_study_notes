package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Tell main that we are adding one goroutine
	wg.Add(1)
	go func() {
		count("sheep", 5)
		// Tell main that defining a goroutine ends here
		wg.Done()
	}()

	go func() {
		count("fish", 200)
		// Tell main that defining a goroutine ends here
		// wg.Done()
	}()
	// with wg.Add(2) and 2 wg.Done(), main() will wait for both goroutine to complete before completes iteself
	// with wg.Add(2) and 1 or 0 wg.Done(), main() will wait for another non-existing goroutine(s) to complete so deadlock occurs.
	// with wg.Add(1) and more than 1 wg.Done(), first gorounties that is marked with Done complete, main() will continue, i.e. in this case completes.
	wg.Wait()
}

func count(thing string, lmt int) {
	for i := 0; i < lmt; i++ {
		fmt.Println(i, thing)
	}
}
