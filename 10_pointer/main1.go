package main

import "fmt"

func main() {
	var a int = 5
	// Assign a's address to b - b is a pointer
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(a, *b)

	c := [3]int{1, 2, 3}
	d := &c[0]
	e := &c[1]
	// difference between d and e is 8 bytes, i.e. 64 bit
	// pointer manipulation is not allowed, i.eg d = e -8 is not allowed
	fmt.Printf("%v %p %p\n", c, d, e)

	// map and slice is a obj with internal pointer, i.e. refer type
}