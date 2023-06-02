package main

import (
	"fmt"
	"reflect"
)

func main() {

	var s []string
	fmt.Println("uninit", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), cap(s))

	t := []string{"a", "b", "c"}

	r := [3]string{"a", "b", "c"}

	var q [3]string

	fmt.Println("emp: s", s, "len:", len(s), cap(s))
	fmt.Println("emp: q", q, "len:", len(q), cap(q))
	v := reflect.ValueOf(q).Kind()
	fmt.Println("reflect.ValueOf: var q [3]string ", v)
	v = reflect.ValueOf(s).Kind()
	fmt.Println("reflect.ValueOf: var s []string", v)
	v = reflect.ValueOf(t).Kind()
	fmt.Println("reflect.ValueOf: t := []string{\"a\", \"b\", \"c\"}", v)
	v = reflect.ValueOf(r).Kind()
	fmt.Println("reflect.ValueOf: r := [3]string{\"a\", \"b\", \"c\"}", v)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[1])

	fmt.Println("len: ", len(s))

	s = append(s, "d", "e", "f")
	fmt.Println("apd: ", s)
	fmt.Println("s:", s, "len:", len(s), cap(s))

	c := make([]string, len(s))
	copy(c, t)

	fmt.Println("cpy: ", c)

	copy(c, s)

	fmt.Println("cpy: ", c)

	l := c[2:5]
	fmt.Println("sl1: ", l)

	m := c[3:]
	fmt.Println("sl1: ", m)

	n := c[:3]
	fmt.Println("sl1: ", n)

	twoD := [][]int{{1, 2}, {3, 4, 5, 6}, {7, 8, 9}}
	fmt.Println("twod: ", twoD)
	threeD := make([][][]int, 10)
	for i := 0; i < 3; i++ {
		innerDimension := i + 1
		threeD[i] = make([][]int, innerDimension)
		for j := 0; j < innerDimension; j++ {
			mostInnerDimension := j + 1
			threeD[i][j] = make([]int, mostInnerDimension)
			for k := 0; k < mostInnerDimension; k++ {
				threeD[i][j][k] = i + j + k
			}
		}
	}
	fmt.Println("3d: ", threeD)

	o := c[2:3]
	fmt.Println("o: ", o, " c:", c)
	o[0] = "reference changed;"
	fmt.Println("o: ", o, " c:", c)

	p := o

	o = c[:cap(c)]
	fmt.Println("o: ", o, " c:", c)

	p[0] = "what"
	fmt.Println("o: ", o, " c:", c)

	z := c

	fmt.Println("z: ", z)

	z = z[2:4]

	fmt.Println("z: ", z)

	z = z[:cap(z)]

	fmt.Println("z: ", z)
	z[0] = "test"
	fmt.Println("z: ", z, "c: ", c)

	ab := make([]int, 2, 2)

	ab = []int{1, 2}

	fmt.Println("ab: ", ab)

	ab = AppendInt(ab, 3, 4)

	fmt.Println("ab: ", ab)
}

func AppendInt(slice []int, data ...int) []int {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]int, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	fmt.Println("slice: ", slice, " len: ", len(slice), " cap: ", cap(slice))
	slice = slice[0:n]
	fmt.Println("slice: ", slice, " len: ", len(slice), " cap: ", cap(slice))
	copy(slice[m:n], data)
	fmt.Println("slice: ", slice, " len: ", len(slice), " cap: ", cap(slice))
	return slice
}
