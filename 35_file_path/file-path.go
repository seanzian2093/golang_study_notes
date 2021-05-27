package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename.ext")
	fmt.Println("p:", p)

	// join will normalize paths by removing superfluous separators and directory changes
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/.../dir1", "filename"))

	// Dir and Base
	fmt.Println("Dir(p): ", filepath.Dir(p))
	fmt.Println("Base(p): ", filepath.Base(p))

	// check if abs path
	fmt.Println("dir/file is abs path: ", filepath.IsAbs("dir/file"))
	fmt.Println("/dir/file is abs path: ", filepath.IsAbs("/dir/file"))

	// split extension
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	// use strings.TrimSuffix to trim ext from filename
	fmt.Println(strings.TrimSuffix(filename, ext))

	// rel finds a relative path of a target to a base
	// throw error if it can not make a relative path
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

}