package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["key1"] = 4
	m["key2"] = 8

	fmt.Println("map: ", m)

	value1 := m["key1"]
	fmt.Println("value1: ", value1)
	value2 := m["key2"]
	fmt.Println("value2: ", value2)

	fmt.Println("len: ", len(m))

	for key, value := range m {

		fmt.Println("key: ", key, " value: ", value)

	}

	delete(m, "key2")
	fmt.Println("map: ", m)

	valueOfKeyTwo, prs := m["key2"]
	fmt.Println("key2-value: ", valueOfKeyTwo, " present: ", prs)

	n := map[string]int{"what": 2, "shall": 3, "we": 4}
	fmt.Println("map2: ", n)

}
