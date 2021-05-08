package main

import "fmt"

// global var


func main() {
	// string
	// bool
	// int
	// int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// float32 float64

	// var keyword
	// declare name as tring and assign a value to it
	var firstname string = "Sean"

	// go can infer type
	var lastname = "Zian"
	var age int32 = 37

	// const keyword
	const isCool bool = true

	fmt.Println(firstname, lastname)
	fmt.Printf("%T\n", firstname)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)

	// for var, we can reassign a value
	age = 32
	fmt.Printf("%v\n", age)

	// for const, we cant reassign
	// isCool = false
	fmt.Printf("%v\n", isCool)

	// shorthand declare and assign
	course := "02_vars"
	fmt.Printf("%v\n", course)

	// assign multiple var
	title, email := "data scientist", "seanzian@gmail.com"
	fmt.Println(title, email)
}
