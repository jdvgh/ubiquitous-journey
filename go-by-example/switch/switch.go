package main

import (
	"fmt"

	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weeeeeekend")
	case time.Friday:
		fmt.Println("It's almost weekend")
	default:
		fmt.Println("It's a weekday")
	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("dunno", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
