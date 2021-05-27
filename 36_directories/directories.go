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

	// similar to mkdir
	err := os.Mkdir("subdir", 0755)
	check(err)
	// similar to rm -rf
	defer os.RemoveAll("subdir")

	// helper func for creating empty file
	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}
	createEmptyFile("subdir/file1")

	// similar to mkdir -p
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ioutil.ReadDir returns a list of os.FileInfo obj
	c, err := ioutil.ReadDir("subdir/parent")
	check(err)
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// os.Chdir similar to cd
	err = os.Chdir("subdir/parent/child")
	check(err)

	c, err = ioutil.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	// filepath.Walk walk through a directory recursively
	// accepts a callback func to handle every file or dir walked, i.e. recursively
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}