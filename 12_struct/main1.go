package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number int
	actorName string
	companions []string
}

func main() {
	// Best practice - pass value by name instead of order
	aDoctor := Doctor{
		number: 3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.number)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[0])

	// Anonimous struct
	bDoctor := struct{name string}{name: "John Hancock"}
	fmt.Println(bDoctor)

	// struct value type, passing by value
	cDoctor := bDoctor
	cDoctor.name = "John Smith"
	fmt.Printf("bDoctor is still: %v\n", bDoctor)
	fmt.Printf("cDoctor is changed to: %v\n", cDoctor)

	// Use pointer to pass by refrence
	dDoctor := &bDoctor
	dDoctor.name = "John Donovan"
	fmt.Printf("bDoctor is changed to: %v\n", bDoctor)
	fmt.Printf("dDoctor is changed to: %v\n", dDoctor)

	// embedding
	type Person struct {
		age int
		gender string
	}

	// A struct uses Person as its field
	type Scientist struct {
		Person
		Domain string
		IsEvil bool
	}

	sct := Scientist{
		Person: Person{age: 36, gender: "male"},
		Domain: "Data and analytisc",
		IsEvil: false,
	}

	// Though Scientis does not have direct field of age, gender, but its Person field has.
	// can access them directly as if they were regular field.
	fmt.Println(sct.age)
	fmt.Println(sct.gender)
	fmt.Println(sct.Domain)
	fmt.Println(sct.IsEvil)

	// Tag - use backtick after type
	type Animal struct {
		Name string `required max: "100"`
		Origin string
	}

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

}