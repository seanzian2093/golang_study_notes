package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}


func main() {

	// easiest way is to use ioutil.TempFile()
	// creates a file for reading and writing and returns a *os.File obj
	// 1st arg set to "" so the tmp file will be created at default dir in os
	// 2nd arg is used as prefix of random filename
	f, err := ioutil.TempFile("", "mytmpfile")
	fmt.Printf("f is of type: %T\n", f)
	check(err)

	// view its name
	fmt.Println("Tmp file name: ", f.Name())
	// remeber to delete it
	defer os.Remove(f.Name())

	// write bytes to the file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// ioutil.TempDir() works similarly but return a dir name
	dname, err := ioutil.TempDir("", "mytmpdir")
	check(err)
	fmt.Println("Tmp dir name: ", dname)

	defer os.RemoveAll(dname)

	// write some to the tmp file
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}