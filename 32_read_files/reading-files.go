package main

import (
	"bufio"
	"fmt"
	"io"
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
	// always close the file when done
	defer f.Close()

	// read from os.File
	b1 := make([]byte, 5)
	// read upto len(b1) and send to buffer b1? though return the number of bytes that read
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// read from a known location
	// Seek sets the offset for the next Read or Write on file to offset, interpreted according to whence:
	// 0 means relative to the origin of the file, 1 means relative to the current offset, and 2 means relative to the end. 
	// It returns the new offset and an error, if any.
	o2, err := f.Seek(6, 0)
	// a new offset, an int64
	fmt.Printf("o2 is of type: %T\n", o2)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d\n", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)

	// io pkg provides i/o utilities too

	// ReadAtLeast reads from r into buf until it has read at least min bytes.
	// It returns the number of bytes copied and an error if fewer bytes were read.
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %v\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	// bufio pkg provides buffed i/o utilities
	r4 := bufio.NewReader(f)
	// Peek returns the next n bytes without advancing the reader.
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

}