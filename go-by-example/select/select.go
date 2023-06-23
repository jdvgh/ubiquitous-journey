package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan bool, 1)
	channel2 := make(chan bool, 1)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- true
	}()
	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- true
	}()
	for i := 0; i < 2; i++ {
		select {
		case ok1 := <-channel1:
			fmt.Printf("Channel 1: %v\n", ok1)
		case ok2 := <-channel2:
			fmt.Printf("Channel 2: %v\n", ok2)

		}
	}
}
