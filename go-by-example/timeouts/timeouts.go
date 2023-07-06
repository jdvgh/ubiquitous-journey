package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan bool, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- true
	}()

	select {
	case ok := <-c1:
		fmt.Printf("c1: %v\n", ok)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}

}
