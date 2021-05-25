package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/create_file.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("Creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing")
	// Fprintf formats according to a format specifier and writes to w. It returns the number of bytes written and any write error encountered
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing")
	err := f.Close()
	if err != nil {
		// os.Stderr, os.Stdout
		fmt.Fprintln(os.Stderr, "error %v\n", err)
		os.Exit(1)
	}
}