package main

import "fmt"

// Closure - a func that returns a func
// adder is a closure, takes no argument ,return a func which takes an int, returns an int
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := adder()
	for i :=0; i <= 10; i++ {
		fmt.Println(sum(i))
	}
}
