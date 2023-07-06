package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(time.Second)

	go func() {
		<-t1.C
		fmt.Println("timer1 fired")
	}()

	t2 := time.NewTimer(2 * time.Second)

	go func() {
		<-t2.C
		fmt.Println("timer2 fired")
	}()

	if ok := t2.Stop(); ok {
		fmt.Println("timer2 stopped")
	}
	time.Sleep(2 * time.Second)

}
