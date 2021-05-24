package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings after sorting: ", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Integers after soring: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Integers are sorted: ", s)
}
