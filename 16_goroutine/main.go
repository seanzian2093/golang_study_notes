package main

import (
	"fmt"
	"sync"
	"runtime"
	"time"
)

// wait groups of goroutines to complete
var wg = sync.WaitGroup{}
var counter = 0
// Mutex to protect data?
var m = sync.RWMutex{}

func main() {
	if false {
		// goroutine - abstract of thread, lightweighter than OS thread
		go sayHello()
		// let main() stay alive a bit while for the goroutine to print
		time.Sleep(100 * time.Millisecond)
	}

	if false {
		var msg = "Hello"
		go func() {
			// func in go routine has access to msg
			// But no guarantee that main() will wait for go routine
			// So here, main() goes on and msg is changed to Goodbye
			// so go routine may print Goodbye
			fmt.Println(msg)
		}()
		msg = "Goodbye"
		time.Sleep(100 * time.Millisecond)
	}

	if false {
		var msg = "Hello"
		// safer way is to pass msg as arg to func in goroutine- pass by value
		go func(msg string) {
			fmt.Println(msg)
		}(msg)
		msg = "Goodbye"
		time.Sleep(100 * time.Millisecond)
	}

	if false {
		var msg = "Hello"
		// add one goroutine to wait group
		wg.Add(1)
		go func(msg string) {
			fmt.Println(msg)
			// in go routine,tell wait group that it is done when it is done
			wg.Done()
		}(msg)
		msg = "Goodbye"
		// let wait group wait as needed
		wg.Wait()
	}

	if true {
		// see available threads- # of CPU cores by default 
		fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
		// set thread number to 1
		// runtime.GOMAXPROCS(100)
		for i := 0; i < 10; i++ {
			wg.Add(2)
			m.RLock()
			go sayHello_1()
			m.Lock()
			go increment()
		}
		wg.Wait()
	}
}

func sayHello() {
	fmt.Println("Hello")
}

func sayHello_1() {
	// m.RLock()
	fmt.Printf("Hello %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	// m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}
