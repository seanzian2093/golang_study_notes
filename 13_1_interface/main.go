package main

import (
	"fmt"
)

func main() {
	// Declare a interface type, and instantiate to a ConsoleWriter
	var w Writer = ConsoleWriter{}
	// Interface's Write method will automatically call ConsoleWriter's Write method
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

// interface is a type
// interface contains behaviors, not data
type Writer interface {
	// Write accepts a byte and return int and/or error
	Write([]byte) (int, error)
}

// Other types can also be implemented in interface
type ConsoleWriter struct {

}
// Define a method for ConsoleWriter - Write
// so a ConsoleWriter w instance can do a w.Write()
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}


type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}