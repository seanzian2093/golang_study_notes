package myownsqrt

import "math"

// if return a value, shall declare its type in func statement
// function name starts with an uppercase to export it, i.e. used in other packages.
func My_sqrt(f float64) float64 {
	var res float64 = math.Sqrt(f)
	return res
}