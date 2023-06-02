package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	for num2 := -5; num2 <= 10; num2++ {
		if num2 < 0 {
			fmt.Println(num2, "is negative")
		} else if num2 < 10 {
			fmt.Println(num2, "has 1 digit")
		} else {
			fmt.Println(num2, "has multiple digits")
		}
	}
}
