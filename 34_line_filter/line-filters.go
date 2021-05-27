package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {

	// Wrapping the unbuffered os.Stdin with a buffered scanner gives us a convenient Scan method that advances the scanner to the next token
	// which is the next line in the default scanner
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Text() returns the current token, i.e. next line from input
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// check error; end of file is expected and not reported as error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
// cat /tmp/lines | go run line-filters.go