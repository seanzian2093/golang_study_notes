package main

import (
	"fmt"
	"strconv"
)
// Define person struct
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// Method #1 - a func attached to a Person instance - value receiver
func (p Person) greet() string {
	// p.age changes at local
	p.age++
	return "Hello, my name is " + p.firstName + ". I am " + strconv.Itoa(p.age) + " years old."
}

// Method #2 - a func attached to a Person instance - pointer receiver, side effect
func (p *Person) hasBirthday() int {
	// p.age changes at address level
	p.age++
	return p.age
}

// Method #3 - a func attached to a Person instance - pointer receiver, side effect
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init a person struct - verbose way
	person1 := Person{firstName: "Sean", lastName: "Zian", city: "Shanghai", gender: "male", age: 29}
	// simple way
	person2 := Person{"Lucia", "Ma", "Shanghai", "female", 24}

	fmt.Println(person1)
	fmt.Println(person2)

	// Access field
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1.age)

	// Access method #1
	fmt.Println(person1.greet())
	fmt.Println(person1.age)

	// Access method #2 - can not be used as a value if it does not return a value; use its side effect
	// person1.hasBirthday()
	fmt.Println(person1.hasBirthday())
	fmt.Println(person1.age)

	// Access method #3
	fmt.Println(person2.lastName)
	person2.getMarried("Zian")
	fmt.Println(person2.lastName)

	// Access method #3
	fmt.Println(person1.lastName)
	person2.getMarried("Ma")
	fmt.Println(person1.lastName)

}
