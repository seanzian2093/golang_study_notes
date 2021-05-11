package main

import (
	"fmt"
	// "math"
)

const a int16 = 27
func main() {
	// name convention
	// camelCase for internal constants - myConstant
	// PascalCase for exported - MyConstant

	// immutable

	// has to be calculable at compile time
	// in below case, math.Sin() is not calculable at compile time, so the initialization fails
	// const myConst float64 = math.Sin(1.57)

	// primitive to be assigned to const - int, float, string, bool
	const myConst float64 = 1.57
	fmt.Printf("%v, %T\n", myConst, myConst)

	// shadowing is possible
	fmt.Printf("The constant a at package level is %v, %T\n", a, a)
	const a int = 99
	fmt.Printf("The constant a at block level is %v, %T\n", a, a)

	// infer the type
	// in below case, the addition operation is essentially 42 + c, and 42 is automatically convert to int32, i.e. b's type, since 42's type is not defined
	const b = 42
	var c int32 = 27
	fmt.Printf("%v, %T\n", b + c, b + c)

	// enumerator const - iota, block level, i.e. reset to 0 in a new block
	const (
		// ignore 1st value by assigning it to a blank identifier
		_ = iota
		d = iota
		e = iota
		f = iota
	)

	const (
		// redefine the starting value
		_ = iota + 7
		g
		h
		i
		j
		k
	)

	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", g, g)
	fmt.Printf("%v, %T\n", h, h)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)
}