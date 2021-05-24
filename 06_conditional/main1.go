package main

import (
	"fmt"
)

func main() {

	statePopulations := map[string]int {
		"California": 39250017,
		"Texas": 27862596,
		"Florida": 2349875,
		"New York": 542908731,
		"Pennsylvania": 93534592,
		"Illinos": 39254351,
		"Ohio": 39897451,
	}
	// Initialize the condition test - if initialier; boolean result {}
	// pop and ok and block level, in within the if scope
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}
	// undefined error
	// fmt.Println(ok)
	// fmt.Println(pop)

	// || or, && and, ! not
	// short circurting is implemented in go
	fmt.Printf("true && false is: %v\n", true && false)
	fmt.Printf("true || false is: %v\n", true || false)
	fmt.Printf("!true is: %v\n", !true)

	// switch -
	// when test eqaul, in each case statement, could be multiple targets but no duplicate allowed
	
	switch 9 {
	case 1, 5, 10:
		// break is implement implicitly
		fmt.Println("one, five, ten")
		// but could be used explicitly
		// break
		// fallthrough will execute the statement in next case no matter condition is true or false
		// fallthrough
		fmt.Println("five, ten")
	case 2, 4, 6:
		fmt.Println("two, four, six")
	default:
		fmt.Println("another number")
	}
}