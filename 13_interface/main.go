package main

import (
	"fmt"
	"math"
)

// Define interface - special struct ?
type Shape interface {
	area() float64
}

// Circle
type Circle struct {
	x, y, radius float64
}

// Rectangle
type Rectangle struct {
	width, height float64
}

// Circle's method - area()
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Rectangle's method - area()
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Shape's method - getArea()
func getArea(s Shape) float64 {
	return s.area()
}


func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
