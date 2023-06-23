package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTicker(200 * time.Millisecond)
	done := make(chan bool)
	go func() {
		counter := 0
		for {
			select {
			case tick := <-t1.C:
				counter++
				fmt.Printf("t1 fired; %v. time at %v\n", counter, tick)
			case <-done:
				fmt.Println("done")
				return
			}
		}
	}()

	time.Sleep(time.Second)
	t1.Stop()
	done <- true
	close(done)
}
