package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	area := r.height * r.width
	return area
}

func (r Rectangle) perimeter() float64 {
	perimeter := 2*r.height + 2*r.width
	return perimeter
}

func (c Circle) area() float64 {
	area := math.Pi * c.radius * c.radius
	return area
}

func (c Circle) perimeter() float64 {
	perimeter := 2 * math.Pi * c.radius
	return perimeter
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	circle := Circle{5}
	square := Rectangle{4, 10}

	measure(circle)
	measure(square)
}
