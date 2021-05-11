package main

import (
	"fmt"
)

func main() {
	// if two iterators - i, j = i++, j++ is not working, because i++ is statement, not expression in go
	// here i and j are within for loop block/scope
	for i, j := 0, 2; i < 5 && j < 5; i, j = i + 1, j + 1{
		fmt.Println(i, j)
	}
	// undefined error will occur
	// fmt.Println(i, j)

	// there is not WHILE keyword in go; but could be constructed like below
	k := 0
	// either 0 ; or 2 ;
	// for k < 100 {
	for ; k < 100 ; {
		fmt.Println(k)
		k = k + 10
	}

	// infinite loops - just drop the terminating condition
	m := 0
	for {
		fmt.Println(m)
		m++
		if m == 9 {
			// break is to leave entire loop
			break
			// continue is to leave current iteration
		}
	}
}