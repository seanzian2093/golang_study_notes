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

func main() {
	fmt.Println(greeting("Sean"))
}