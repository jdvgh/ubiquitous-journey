package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{name: name, age: age}
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Bobby", age: 200})
	fmt.Println(person{name: "Bobsters"})
	fmt.Println(person{age: 420})

	newp := newPerson("Henry", 67)
	fmt.Println("newPerson: ", newp)

	fmt.Println("&newPerson: ", &newp)

	s := person{name: "Dean", age: 60}
	fmt.Println("s.name: ", s.name)

	sp := &s
	fmt.Println("sp.age: ", sp.age)

	sp.age = 51
	fmt.Println("sp.age: ", sp.age)
	fmt.Println("s.age: ", s.age)

	animal := struct {
		name   string
		isGood bool
	}{
		"Peter",
		true,
	}
	fmt.Println("animal: ", animal)

}
