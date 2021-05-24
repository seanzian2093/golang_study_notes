package main

import "fmt"

// one argument, of type string ,return string
func greeting(name string) string {
	return "Hello" + name
}

// two arguments
func getSum(num1 int, num2 int) int {
	return num1 + num2
}

// recursion
func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num - 1)
}

func main() {
	fmt.Println(greeting("Sean"))
	fmt.Printf("fact(9) is %v\n", fact(9))
}