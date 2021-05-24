package main

import (
	"fmt"
	"sort"
)

// define own type 
// to use sort package's generic Sort function Len, Less and Swap methods must be implemented by interface
type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}


func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	// cast to by Length
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}