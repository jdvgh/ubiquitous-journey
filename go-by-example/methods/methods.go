package main

import (
	"fmt"
	"math"
)

type triangle struct {
	a, b, c float64
}

func (t *triangle) area() float64 {
	// Refer to: https://en.wikipedia.org/wiki/Area_of_a_triangle#Using_side_lengths_(Heron's_formula)
	perimeter := t.perimeter()
	p_a := t.perimeter() - t.a
	p_b := t.perimeter() - t.b
	p_c := t.perimeter() - t.c
	return math.Sqrt(perimeter * p_a * p_b * p_c)

}

func (t triangle) perimeter() float64 {
	return 0.5 * (t.a + t.b + t.c)
}

func main() {
	triangleSameLength := triangle{6.5, 6.5, 6.5}
	fmt.Println("perimeter: ", triangleSameLength.perimeter())
	fmt.Println("area: ", triangleSameLength.area())

	tp := &triangleSameLength
	fmt.Println("perimeter: ", tp.perimeter())
	fmt.Println("area: ", tp.area())
}
