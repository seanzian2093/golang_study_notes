package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 1)
	done := make(chan bool)

	go func() {
		// a infine loop
		for {
			// more will be false if jobs is closed
			j, more := <- jobs
			if more {
				fmt.Println("Job Received: ", j)
			} else {
				fmt.Println("Received all jobs.")
				// send a value to done to notify that all jobs has been worked.
				done <- true
				// return statement to break the infinite loop
				return
			}
		}
	}()

	// add job to jobs
	// synchronically, above worker goroutine will be working on the jobs
	for j :=1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Job sent: ", j)
	}
	// close jobs after sending all jobs, not working
	close(jobs)
	fmt.Println("Sent all jobs.")

	// receiving from done, it is blocking, i.e. main() will proceed after a value is sent done and received
	<-done
}