package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {

    p := point{1, 2}
	// print value
    fmt.Printf("Print the value: %v\n", p)

	// if value is a struct, print its field names
    fmt.Printf("%+v\n", p)

	// print Go syntax representation of the value, i.e. the source code snippet that would produce that value
    fmt.Printf("%#v\n", p)

	// print type
    fmt.Printf("%T\n", p)

    fmt.Printf("%t\n", true)

	// standard, base-10 formatting of integer
    fmt.Printf("%d\n", 123)

	// binary representation
    fmt.Printf("%b\n", 14)

	// character corresponding to the integer
    fmt.Printf("%c\n", 33)

	// hex encoding
    fmt.Printf("%x\n", 456)

	// basic decimal formatting of floats
    fmt.Printf("%f\n", 78.9)

	// scientific notation
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)

	// basic string printing
    fmt.Printf("%s\n", "\"string\"")

	// raw string?
    fmt.Printf("%q\n", "\"string\"")

	// hex representation of string, 
    fmt.Printf("%x\n", "hex this")

	// print a pointer
    fmt.Printf("%p\n", &p)

	// To specify the width of an integer, use a number after the % in the verb. 
	// By default the result will be right-justified and padded with spaces.
    fmt.Printf("|%6d|%6d|\n", 12, 345)

	// similarly to floats
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// use - flog for left justfy
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// width and justfy with string
    fmt.Printf("|%6s|%6s|\n", "foo", "b")

    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// so far Printf() print to os.Stdout
	// Sprintf() formats the string and return it without printing it to anywhere
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)

	// format and print to io.Writers other than os.Stdout
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}