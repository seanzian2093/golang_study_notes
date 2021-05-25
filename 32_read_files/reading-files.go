package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// use ioutil, return byte
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))

	// os.Open, return *os.File
	f, err := os.Open("/tmp/dat")
	fmt.Printf("f is of type: %T\n", f)
	check(err)

	// read from os.File
	b1 := make([]byte, 5)
	// read and send to b1? though return the number of bytes that read
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
}