package main

import (
	"bufio"
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
	// ioutil.WriteFile to dump a string or just bytes into a file
	d1 := []byte("hello\ngo\n")
	// WriteFile writes data to a file named by filename. If the file does not exist, WriteFile creates it with permissions perm (before umask);
	// otherwise WriteFile truncates it before writing, without changing permissions.
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// for more granular writes, open a file for writing
	f, err := os.Create("/tmp/dat2")
	fmt.Printf("f is of type: %T\n", f)
	check(err)

	defer f.Close()

	// write slice of byte to *os.File
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("Wrote %d bytes\n", n2)
	// write string to *os.File
	n3, err := f.WriteString("Write this string\n")
	check(err)
	fmt.Printf("Wrote %d bytes\n", n3)
	
	// Sync commits the current contents of the file to stable storage. 
	//Typically, this means flushing the file system's in-memory copy of recently written data to disk.
	f.Sync()

	// bufio provides buffered writer
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("Wrote %d bytes\n", n4)

	// Use flush to ensure all buffered operations have been applied to the underlying writer
	w.Flush()
}
