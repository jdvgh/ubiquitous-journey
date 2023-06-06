package main

import "fmt"

func main() {

	var nums []int
	nums = []int{2, 3, 4}
	var sum int
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		fmt.Println("index: ", i, " num: ", num)
	}

	kvs := map[string]string{"a": "apple", "b": "bean"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
