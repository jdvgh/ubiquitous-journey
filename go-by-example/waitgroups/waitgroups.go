package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("w%v started\n", id)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("w%v finished\n", id)
}
func main() {

	var wg sync.WaitGroup

	const numWorkers = 5

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)

		i2 := i // i itself is the same reference if we'd use it directly
		go func() {
			defer wg.Done()
			worker(i2)
		}()
	}

	wg.Wait()

}
