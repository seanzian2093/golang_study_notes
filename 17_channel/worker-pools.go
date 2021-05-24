package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker ", id, "started job ", job)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "finished job ", job)
		results <- job * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Initialize 3 workers, operating concurrently
	for w := 1; w <= 3; w++ {
		go worker (w, jobs, results)
	}

	// Initialize 5 jobs and send them to channel jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// collect results by receiving from channel results
	// this ensures all goroutine finished.
	// combination of id & job might be random
	for r := 1; r <= 5; r++ {
		<-results
	}
}