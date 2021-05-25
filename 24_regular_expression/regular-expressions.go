package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {

	// use string pattern directly
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

	// compile an regexp struct
    r, _ := regexp.Compile("p([a-z]+)ch")

	// match test, bool
    fmt.Println(r.MatchString("peach"))

	// find the matched string
    fmt.Println(r.FindString("peach punch"))

	// start and end indexes of FIRST match
    fmt.Println(r.FindStringIndex("peach punch"))

	// return information for both p([a-z]+)ch and ([a-z]+).
	// matched string
    fmt.Println(r.FindStringSubmatch("peach punch"))
	// index of matched string
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// Similar functions with All variant apply to all matches, not just FIRST
    fmt.Println(r.FindAllString("peach punch pinch", -1))

    fmt.Println(r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

    fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Our examples above had string arguments and used names like MatchString. 
	// We can also provide []byte arguments and drop String from the function name, like Match, as opposed to MatchString
    fmt.Println(r.Match([]byte("peach")))

	// MustCompile panics instead of returning an error, which makes it safer to use for global variables.
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)

	// The regexp package can also be used to replace subsets of strings with other values.
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// The Func variant allows you to transform matched text with a given function
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}