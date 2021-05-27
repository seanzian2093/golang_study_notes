package main

import (
    "flag"
    "fmt"
)

func main() {

	// 3 args to flag.*() are name, default value and short description
    wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

	// could also use existing var; need to pass a pointer
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

	// after declaring flags, Parse() them and execute them
    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
	// use flag.Args() to get trailing flags
    fmt.Println("tail:", flag.Args())
}

// compile it and run the binary
// go build cmd-flags.go


// ./cmd-flags -word=opt -numb=7 -fork -svar=flag
// word: opt
// numb: 7
// fork: true
// svar: flag
// tail: []

// Note that the flag package requires all flags to appear before positional arguments
// otherwise the flags will be interpreted as positional arguments).
// ./cmd-flags -word=opt a1 a2 a3 -numb=7
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: [a1 a2 a3 -numb=7]

// If you provide a flag that wasn’t specified to the flag package, the program will print an error message and show the help text again.
// ./cmd-flags -wat
// flag provided but not defined: -wat
// Usage of ./cmd-flags:
//   -fork
//         a bool
//   -numb int
//         an int (default 42)
//   -svar string
//         a string var (default "bar")
//   -word string
//         a string (default "foo")