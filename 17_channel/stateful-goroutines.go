package main

import (
	"fmt"
	"sync/atomic"
	"math/rand"
	"time"
)

// a read operation includes a key to read from, and a resp to indicator sucess; val will be assigned to resp upon success.
type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	// channel to hold readOp and writeOp
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// a stateful goroutine, with an infinite loop to receive from channel
	go func() {
		var state = make(map[int]int)

		for {
			select {
				// when receive from channel reads
			case read := <-reads:
				// update resp
				read.resp <- state[read.key]
			case write := <- writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Initialize 100 goroutines, each contain an infinite loop of readOp
	for r := 1; r <= 100; r++ {
		go func() {
			for {
				read := readOp{key: rand.Intn(5), resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}

		}()
	}

	// Initialize 10 goroutines, each contain an infinite loop of writeOp
	for w := 1; w <= 10; w++ {
		go func() {
			for {
				write := writeOp{key: rand.Intn(5), val: rand.Intn(100), resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}

		}()
	}

	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOpsFinal:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOpsFinal:", writeOpsFinal)

}
