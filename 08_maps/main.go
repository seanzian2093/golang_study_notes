package main

import "fmt"

func main() {
	// Define map- key type and value type
	emails := make(map[string]string)

	// Assign keys and values
	emails["Bob"] = "bob@gmail.com"
	emails["Sean"] = "sean@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// Delete from map by key
	delete(emails, "Bob")
	fmt.Println(emails)

	// Declare and assign
	persons := map[string]int{"Bob": 23, "Sean": 29}
	fmt.Println(persons)
}
