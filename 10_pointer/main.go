package main

import "fmt"

func main() {
	a := 5
	// Assign a's address to b - b is a pointer
	b := &a
	fmt.Printf("a= %d, is stored at %d in memory\n", a, b)
	fmt.Printf("a is of type %T, b is of type %T\n", a, b)

	// use * to access value from address - read and write
	fmt.Printf("b's value is %d, the value stored at b is %d\n", b, *b)

	*b = 10
	fmt.Printf("b's value is %d, but now value stored at b is %d\n", b, *b)
	fmt.Printf("Now value of a is %d\n", a)
}
