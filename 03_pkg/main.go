package main

// import multiple packages
import (
	"fmt"
	"math"
	// path must in $GOROOT/src
	// or starts with module name in go.mod
	// the go file name does not matter
	// in the go file, package statment and func statement determine the namespace
	"crash_course/03_pkg/mysqrt"
)
func main() {
	fmt.Println(math.Floor(2.3))
	fmt.Println(math.Ceil(2.3))
	fmt.Println(math.Sqrt(2.3))
	fmt.Println(myownsqrt.My_sqrt(2.3))
}