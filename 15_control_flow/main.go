package main

import (
	"fmt"
	"log"
)

func main() {
	if false {
		fmt.Println("start")
		// defer puts the statement after all statments in main(), just before return values to caller.
		// first defered statement is executed last; verse versa- LIFO, last-in-first-out
		defer fmt.Println("middle")
		fmt.Println("end")

		defer fmt.Println("middle 1")
		defer fmt.Println("middle 2")
		defer fmt.Println("middle 3")

		a := "start"
		// a will be evaluated eagerly, 
		// a is evaluated when defered; not in fmt.Println() is executed
		// i.e. in this defer statment, a is start
		defer fmt.Println(a)
		a = "end"
	}

	// panic
	if false {
		fmt.Println("start")
		// defered statements are executed before paniced statment
		defer fmt.Println("this was defered")
		// throw an error or warning on purpose.
		panic("c might be zero.")
		fmt.Println("end")
		// fmt.Println(ans)
	}

	// recover
	if true {
		fmt.Println("start")
		panicker()
		fmt.Println("end")
	}

}
func panicker() {
	fmt.Println("about to panic")

	defer func() {
		// recover from panics
		// only useful in deferred functions
		// current function will not continue; but higher caller will
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err)
		}
	}()

	panic("something bad happened")
	fmt.Println("done panic")
}