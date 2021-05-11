package main

import (
	"fmt"
)

func main() {
	// Declare a map
	statePopulations := map[string]int {
		"California": 39250017,
		"Texas": 27862596,
		"Florida": 2349875,
		"New York": 542908731,
		"Pennsylvania": 93534592,
		"Illinos": 39254351,
		"Ohio": 39897451,
	}

	fmt.Println(statePopulations)

	// Add a key value pair
	// return order of a map is not guaranteed
	statePopulations["Geogia"] = 10310371
	fmt.Println(statePopulations)

	// Delete an entry
	delete(statePopulations, "Ohio")
	fmt.Println(statePopulations)

	// Cautious - 
	fmt.Printf("After delete Ohio, if access it, we get not error but: %v\n", statePopulations["Ohio"])
	// Deeper investigation - accessing map by key return a tuple
	pop, ok := statePopulations["Ohio"]
	fmt.Println(pop, ok)
	pop1, ok1 := statePopulations["Geogia"]
	fmt.Println(pop1, ok1)


	// map is reference, not value
	// sp and statePopulations refer to one address in memory
	sp := statePopulations
	// delete from sp is essentially from statePopulations too
	delete(sp, "New York")
	fmt.Println(sp)
	fmt.Println(statePopulations)

}