package main

import (
    "fmt"
    "os"
)

// Note that unlike e.g. C, Go does not use an integer return value from main to indicate exit status. 
// If you’d like to exit with a non-zero status you should use os.Exit.
// use os.Exit() to exit immediately with the provided status code
func main() {

    // defer will not be run when using os.Exit()
    defer fmt.Println("!")

    os.Exit(3)
}