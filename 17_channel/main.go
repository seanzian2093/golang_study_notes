package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// Decluare a channel that accepts integer
	// once declared, not able to accept other type data
	// bydefault, no buffer channel, i.e. only one data at a time
	if true {
		ch := make(chan int)
		for j := 0; j < 5; j++ {
			wg.Add(2)
			go func() {
				// receive data from a channel
				i := <- ch
				fmt.Println(i)
				wg.Done()
			}()

			go func() {
				// send data into a channel
				ch <- 42
				wg.Done()
			}()
		}
		wg.Wait()
	}

	if true {
		ch := make(chan int)
		wg.Add(2)

		// Declare this goroutine as recerver only
		go func(ch <-chan int) {
			// receive a data
			i := <- ch
			fmt.Println(i)
			wg.Done()
		}(ch)

		// Declare this goroutine as sender only
		go func(ch chan<- int) {
			// send a data
			ch <- 42
			ch <- 27
			wg.Done()
		}(ch)

		wg.Wait()
	}

	if true {
		// make a channel with buffer of size 50
		ch := make(chan int, 50)
		wg.Add(2)

		// Declare this goroutine as recerver only
		go func(ch <-chan int) {
			// receive a data
			// i := <- ch
			// fmt.Println(i)

			// receive another data
			// i = <- ch
			// fmt.Println(i)

			// or more generic way - use loop over range
			for i := range ch {
				fmt.Println(i)
			}
			wg.Done()
		}(ch)

		// Declare this goroutine as sender only
		go func(ch chan<- int) {
			// send a data
			ch <- 42
			ch <- 27
			// close the channel - otherwise the loop over range would not know when to stop
			close(ch)
			wg.Done()
		}(ch)

		wg.Wait()
	}
}

