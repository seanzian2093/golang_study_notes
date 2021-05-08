package main

import (
	"fmt"
	"strconv"
)

// Declare a variable at package level; has to be full declaration
// variable of name that starts with lower case letter is package level
// with Upper case, it will be exposed to outside of package
var pkg_var int = 99

// varialbes defined within {} is block-level
func main() {
	var i int
	i = 42
	fmt.Println(i)
	// Reassign is okay; re-declare is not
	i = 27
	fmt.Println(i)
	// type will be infered if not declared
	fmt.Printf("%v, %T\n", i, i)
	j := 27.

	fmt.Printf("%v, %T\n", j, j)
	// we can access package variable in a block
	fmt.Printf("%v, %T\n", pkg_var, pkg_var)

	// convert an integer to float
	k := float32(i)
	fmt.Printf("%v, %T\n", k, k)

	// convert a float to integer, decimal part will be lost
	m := 34.5
	l := int(m)
	fmt.Printf("%v, %T\n", l, l)

	// be careful when convert integer to string
	i = 42
	// here unicode of 42 is assign to s, which is a astror *
	// use strconv package instead
	var s string = string(i)
	fmt.Printf("%v, %T\n", s, s)

	// Itoa - integer to ascii string
	var s1 string = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", s1, s1)
}