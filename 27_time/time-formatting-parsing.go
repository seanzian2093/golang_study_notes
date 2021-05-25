package main

import (
    "fmt"
    "time"
)

func main() {
	p := fmt.Println
	t := time.Now()

	// format a time value 
	p(t.Format(time.RFC3339))

	// parse string to get a time value
	// requires a format to tell go how to parse it
	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)

	// Format and Parse use example-based layout
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)
	p(e)

}