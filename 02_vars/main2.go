package main

import (
	"fmt"
)

func main() {
	var n bool = false
	fmt.Printf("%v, %T\n", n, n)
	
	// logic test's result is boolean
	// logic test > assign
	k := 1 == 1
	m := 1 ==2
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", m, m)

	// if a var is not initialized, it automatically to 0
	var j bool
	fmt.Printf("j is initialized at %v, %T\n", j, j)

	// string default to ''
	var s string
	fmt.Printf("s is initialized at %v, %T\n", s, s)

	// int / int yield an int, decimal part is discarded
	var a int = 10
	var b int = 3
	fmt.Printf("a/b is %v\n", a/b)

	// different integers/float have to be converted to same type before calculation 
	var c int = 10
	var d int8 = 3
	fmt.Printf("c/d is %v\n", c/int(d))

	// string is like array, each element is the letter's ascii number of unsigned integer 8
	// , aka byte or utf-8 character
	// string is immutable
	s1 := "This is a string"
	b1 := []byte(s1)
	fmt.Printf("%v, first element is %v of %T\n", s1, string(s1[0]), string(s1[0]))
	fmt.Printf("%v, first element is %v of %T\n", s1, s1[0], s1[0])
	fmt.Printf("%v, %T\n", b1, b1)

	// string should be in double quote
	// rune in single quote, int32, utf-32 character
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

}