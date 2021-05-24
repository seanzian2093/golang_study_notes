package main

import "fmt"

func main() {

	// By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. 
	// Buffered channels accept a limited number of values without a corresponding receiver for those values.

	// make a channel with buffering up to 2 values
    messages := make(chan string, 2)

	// send values to channel without coresponding receivings
    messages <- "buffered"
    messages <- "channel"

	// receive as usual
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}