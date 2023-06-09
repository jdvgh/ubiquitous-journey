package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	fmt.Println(fruits)
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	numberStrings := []string{"1234", "1", "90129301", "1231512"}
	fmt.Println(numberStrings)
	sort.Sort(byLength(numberStrings))
	fmt.Println(numberStrings)

}
