package main

import "fmt"

func main() {

	i := 0

	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j <= 7; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
