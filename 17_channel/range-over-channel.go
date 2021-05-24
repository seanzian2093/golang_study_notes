package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// queue is closed but remaining values are still accessible
	for elem := range queue {
		fmt.Println(elem)
	}
}