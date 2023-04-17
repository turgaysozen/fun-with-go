package main

import (
	"fmt"
	"math"
)

type geo interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2*r.height + 2*r.width
}

func calculate(g geo) {
	fmt.Printf("Geo: %T, \nArea: %v\nPerim: %v\n", g, g.area(), g.perim())
}
