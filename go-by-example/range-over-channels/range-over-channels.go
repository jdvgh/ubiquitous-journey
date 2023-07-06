package main

import (
	"fmt"
	"math/rand"
)

func main() {

	c1 := make(chan float32, 2)
	c1 <- rand.Float32()
	c1 <- rand.Float32()
	close(c1)

	for fl := range c1 {
		fmt.Printf("received on c1: %v\n", fl)
	}
}
