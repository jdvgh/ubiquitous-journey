package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[3])
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a))

	a = [5]int{1, 2, 3, 4, 5}
	fmt.Println("ovrd: ", a)

	b := [2]int{6, 7}
	fmt.Println("dcl: ", b)

	c := [2][3]int{{2, 3, 4}, {5, 6, 7}}
	fmt.Println("dcl 2darray: ", c)

	for i := 0; i <= len(c)-1; i++ {
		for j := 0; j <= len(c[i])-1; j++ {
			fmt.Println("c[", i, "][", j, "]:", c[i][j])
		}
	}
}
