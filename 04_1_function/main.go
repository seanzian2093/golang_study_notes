package main

import "fmt"

// Function name determines its visibility.
// Opening curly brace has to be same line with func keyword
func main() {

	if false {
		for i := 0; i < 1; i++ {
			sayMessage("Hello Go!", i)
		}
	}

	if false {
		greeting := "hello"
		name := "Sean"
		sayGreeting(greeting, name)
		fmt.Printf("In main, the value of name is: %v\n", name)
	}

	if false {
		greeting := "hello"
		name := "Sean"
		sayGreeting_pt(&greeting, &name)
		fmt.Printf("In main, the value of name is: %v\n", name)
	}

	if false {
		s := sum(1, 2, 3, 4, 5)
		fmt.Printf("The sum is: %v\n", s)
	}

	if false {
		s := sum_pt(1, 2, 3, 4, 5)
		// s is a pointer
		fmt.Printf("The pointer of sum is: %v, %T\n", s, s)
		fmt.Printf("The value of sum is: %v, %T\n", *s, *s)
	}

	if false {
		d, err := divide(5.0, 0.0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(d)
	}

	// anonimous function
	if false {
		func() {
			fmt.Println("Hello Go from anonimous function")
		}()
	}
	// function is also a type, can be used as other types are 
	if false {
		var f func() = func() {
			fmt.Println("Hello Go from a var of function type")
		}
		// value is address
		fmt.Printf("%v, %T\n", f, f)
		f()
	}

	// method
	g := greeter{greeting: "Hello", name: "Go"}
	g.greet()
	fmt.Println("The new name is:", g.name)
// Closing curly brace be on its own line.
}

// No need to use var key word for setting up arguments.
func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

// go will infer that both greeting and name are both string
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
	// change name's value
	name = "Lucia"
	fmt.Printf("In sayGreeting, after change, the value of name is: %v\n", name)
}

// This func accept pointers of string var
// map and slice are with internal pointers so no need to use this approach.
func sayGreeting_pt(greeting, name *string) {
	fmt.Println(*greeting, *name)
	// change name's value
	*name = "Lucia"
	fmt.Printf("In sayGreeting_pt, after change, the value of name is: %v\n", *name)
}


// variomatic variables - use ... , similar to * in python
func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}

	return result
}

// function can also return a point
// define the return type is address of int, instead of int
func sum_pt(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}

	// return the address instead of value
	return &result
}

// func can return multiple values
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by 0")
	}

	return a / b, nil
}

// method
type greeter struct {
	greeting string
	name string
}
// this special func, a method run in the context of g, a greeter type
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "Sean"
}