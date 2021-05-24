package main

import (
	"fmt"
)

func main() {
	// In below declaration, size is declared twice - in [] and {}
	grades := [3]int{97, 85, 93}
	// Use [...] to create a size that just suffice
	names := [...]string{"Sean", "Lucia", "Emma"}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Names: %v\n", names)

	// Declare without assign, must use full declare method, not := 
	var courses [3]string
	fmt.Printf("Course: %v\n", courses)
	courses[0] = "Math"
	courses[1] = "English"
	courses[2] = "Physics"
	fmt.Printf("Course: %v\n", courses)

	// multiple dim array, i.e. array of arrays
	// a 3 * 3 matrix of integers
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// array is value type, i.e. pass by value, not refer
	arr1 := [...]int{1, 2, 3}
	// value of arr1 is copied to arr2
	arr2 := arr1
	// now changing arr2 does not apply to arr1
	arr2[0] = 99
	fmt.Println(arr1)
	fmt.Println(arr2)

	// use pointer for passing by refer
	arr3 := &arr1
	arr3[0] = 999
	fmt.Println(arr1)
	fmt.Println(arr3)

	// slice - array with predefined size
	slc1 := []int{1, 2, 3}
	fmt.Printf("Length: %v\n", len(slc1))
	fmt.Printf("Capacity: %v\n", cap(slc1))
	// slice is reference type, pass by reference, NOT value
	slc2 := slc1
	// changes in slc2 is changes in slc1
	slc2[0] = 88
	fmt.Println(slc1)
	fmt.Println(slc2)

	// slice
	slc3:= []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slc3[:]) // all elements
	fmt.Println(slc3[3:]) // 4th to end
	fmt.Println(slc3[:6]) // 1st to 6th
	fmt.Println(slc3[3:6]) // 4th to 6th

	// make(type, length, capacity) - if capacity if present, default to length
	make1 := make([]int, 3)
	fmt.Println(make1)
	fmt.Printf("Length of make1: %v\n", len(make1))
	fmt.Printf("Capacity of make1: %v\n", cap(make1))

	make2 := make([]int, 3, 100)
	fmt.Println(make2)
	fmt.Printf("Length of make2: %v\n", len(make2))
	fmt.Printf("Capacity of make2: %v\n", cap(make2))

	// append()
	apd1 := []int{}
	fmt.Println(apd1)
	fmt.Printf("Length of apd1: %v\n", len(apd1))
	fmt.Printf("Capacity of apd1: %v\n", cap(apd1))

	// append return a new one, i.e. copy, apd1 is not changed
	// all args after 1st is appended to 1st arg
	apd2 := append(apd1, 2, 3, 4)
	fmt.Println(apd1)
	fmt.Printf("After append, length of apd1: %v\n", len(apd1))
	fmt.Printf("After append, capacity of apd1: %v\n", cap(apd1))

	fmt.Println(apd2)
	fmt.Printf("Length of apd2: %v\n", len(apd2))
	fmt.Printf("Capacity of apd2: %v\n", cap(apd2))

	// unpacking or spread - passing a slice directly to append is not allowed
	// apd3 := append(apd1, []int{2, 3, 4})
	// must unpack it by trailing ...
	apd3 := append(apd1, []int{9, 99, 999}...)
	fmt.Println(apd3)
	fmt.Printf("Length of apd3: %v\n", len(apd3))
	fmt.Printf("Capacity of apd3: %v\n", cap(apd3))





}