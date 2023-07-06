package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan float32, 5)
	done := make(chan bool)

	go func() {
		for {
			signal, more := <-c1
			if more {
				fmt.Printf("but wait, there's more - another float32: %v\n", signal)
			} else {
				fmt.Println("stream closed")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 5; i++ {
		var randFloat = rand.Float32()
		c1 <- randFloat
		fmt.Printf("Sent: %v \n", randFloat)
	}
	close(c1)
	fmt.Println("Closed float channel")
	time.Sleep(1 * time.Second) // This or the done statement is needed to await on the worker, otherwise we close before the goroutine notices
	<-done

}
