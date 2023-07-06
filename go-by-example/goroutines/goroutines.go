package main

import (
	"fmt"
	"time"
)

func iterator(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, " : ", i)
	}
}

func main() {
	iterator("direct")

	go iterator("goroutine")

	go func(message string) {
		fmt.Println(message)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

}
