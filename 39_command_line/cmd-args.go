package main

import (
    "fmt"
    "os"
)

func main() {

	// os.Args is a slice of araw arguments
	// 1st is the path to the program; 2nd and beyond is args to the command
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}

// compile it and run the binary
// go build cmd-args.go
// ./cmd-args a b c d