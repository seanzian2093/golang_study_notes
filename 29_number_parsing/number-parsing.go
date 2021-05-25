package main

import (
    "fmt"
    "strconv"
)

func main() {

	// convert a string to float, 64 tells how many bits of precision to parse
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

	// convert a string to integer, 0 means infer the base, 64 mean precision
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

	// hex coded are automatically recognized
    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

	// to unsigned integer
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

	// a convenient converter from string to 10-base integer
    k, _ := strconv.Atoi("135")
    fmt.Println(k)

	// Parse return error on bad input
    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}