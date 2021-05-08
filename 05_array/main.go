package main

import "fmt"

func main() {
	// Declare an array, len=2, of 2 elements
	var fruitArr [2]string

	// Assign values, 0-based
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Or declare and assign - length,type and elements
	names := [2]string{"Sean", "Zian"}

	// Slice
	fruitSlice := []string{"Apple", "Orange", "Grape"}



	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(names)
	fmt.Println(fruitSlice)
	// starts at 1 stops at2, exclusive
	fmt.Println(fruitSlice[1:2])
}
