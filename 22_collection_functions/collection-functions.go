package main

import (
	"fmt"
	"strings"
)

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Inclue(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool{return strings.HasPrefix(v, "p")}))

	fmt.Println(All(strs, func(v string) bool{return strings.HasPrefix(v, "p")}))

	fmt.Println(Filter(strs, func(v string) bool{return strings.Contains(v, "e")}))

	fmt.Println(Map(strs, strings.ToUpper))
}

func Index(vs []string, s string) int {
	for i, v := range vs {
		if v == s {
			return i
		}
	}
	// return -1 if not found
	return -1
}

func Inclue(vs []string, s string) bool {
	return Index(vs, s) >= 0
}

// here func(string) bool is working as a type
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}

	return vsm
}