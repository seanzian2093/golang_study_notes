package main

import (
	"fmt"
	"time"
)

func main() {
	// a channel of string, with buffer 1
	// buffered channel is nonblocking; while basic send and receive is blocking

	// after sleeping 2 sec, send a message to c1
	c1 := make(chan string, 1)
	go func () {
		time.Sleep(2 * time.Second)
		c1 <- "Result 1"
	}()

	// meanwhile after 1 second, time out message is sent to case
	select {
	case res := <- c1:
		fmt.Println(res)
	case <- time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// after sleeping 2 sec, send a message to c2
	c2 := make(chan string, 1)
	go func () {
		time.Sleep(2 * time.Second)
		c2 <- "Result 2"
	}()

	// meanwhile after 3 second, time out message will be sent to case which will NOT happen because after 2 sec, c2 received the message
	// so the 1st case will be selected
	select {
	case res := <- c2:
		fmt.Println(res)
	case <- time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}