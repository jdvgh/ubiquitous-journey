package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type triangle struct {
	a, b, c float64
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

func (t triangle) area() float64 {
	// Refer to: https://en.wikipedia.org/wiki/Area_of_a_triangle#Using_side_lengths_(Heron's_formula)
	perimeter := t.perim()
	p_a := t.perim() - t.a
	p_b := t.perim() - t.b
	p_c := t.perim() - t.c
	return math.Sqrt(perimeter * p_a * p_b * p_c)

}

func (t triangle) perim() float64 {
	return 0.5 * (t.a + t.b + t.c)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	t := triangle{5.5, 5.5, 5.5}
	c := circle{32}

	measure(t)
	measure(c)

}
