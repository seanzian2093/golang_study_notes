package main

import "fmt"

func main() {
	// Declare a slice and assign
	ids := []int{33,76,54,23,11,2}

	// Loop through ids- range yield a (index, value) pairs
	for i, id :=range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Loop through ids- not using index
	for _, id :=range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id :=range ids {
		sum += id
	}
	fmt.Printf("Sum: %d\n", sum)

	// Range with map - yield (key, value) pairs
	persons := map[string]int{"Bob": 23, "Sean": 29, "Lucia": 18}
	for k, v := range persons {
		fmt.Printf("%s is %d years old.\n", k, v)
	}

	// Range with map - use one
	for k := range persons {
		fmt.Printf("Name is %s.\n", k)
	}
}
