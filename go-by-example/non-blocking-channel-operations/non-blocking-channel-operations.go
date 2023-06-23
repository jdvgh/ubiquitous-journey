package main

import "fmt"

func main() {

	c1 := make(chan bool, 1)
	c2 := make(chan bool, 1)

	select {
	case ok := <-c1:
		fmt.Printf("c1: %v\n", ok)
	default:
		fmt.Println("no signal on c1")
	}

	c2 <- true
	select {
	case ok := <-c2:
		fmt.Printf("c2: %v\n", ok)
	default:
		fmt.Println("no signal on c2")
	}

	select {
	case ok := <-c1:
		fmt.Printf("c1: %v\n", ok)
	case ok := <-c2:
		fmt.Printf("c2: %v\n", ok)
	default:
		fmt.Println("no signal at all")
	}

}
