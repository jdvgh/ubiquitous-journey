package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	fmt.Println("Strings: ", strs)
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

	ints := []int{7, 2, 4}
	fmt.Println("Ints: ", ints)
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted ints: ", s)

	r := sort.StringsAreSorted(strs)
	fmt.Println("Sorted strings: ", r)
}
