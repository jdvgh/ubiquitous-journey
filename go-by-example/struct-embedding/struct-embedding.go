package main

import "fmt"

type printable interface {
	toString()
}
type describer interface {
	describe() string
}
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v\n", b.num)
}

func (b base) toString() {
	fmt.Printf(b.describe())
}

type container struct {
	base
	str string
}

func (c container) toString() {
	fmt.Printf("c={num: %v, str: %v}\n", c.num, c.str)
	fmt.Println("also num:", c.base.num)
	c.base.toString()
	fmt.Printf(c.base.describe())
	fmt.Printf(c.describe())
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "container 1",
	}
	co2 := container{
		base{
			num: 2,
		},
		"container 2",
	}
	co.toString()
	co2.toString()

	var d1 describer = co
	var d2 describer = co2

	fmt.Println("describer:", d1.describe())
	fmt.Println("describer:", d2.describe())

}
