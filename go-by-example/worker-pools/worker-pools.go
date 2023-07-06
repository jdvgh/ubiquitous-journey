package main

import (
	"fmt"
	"time"
)

func worker(jobid int, tasks <-chan int, output chan<- int) {
	for task := range tasks {
		fmt.Printf("w: %v, task: %v - started \n", jobid, task)
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("w: %v, task: %v - finished \n", jobid, task)
		output <- jobid + task
	}

}

func main() {
	const numTasks = 5
	tasks := make(chan int, numTasks)
	outputs := make(chan int, numTasks)

	for i := 10; i < 60; i = i + 10 {
		go worker(i, tasks, outputs)
	}

	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}
	close(tasks)

	for out := 1; out <= numTasks; out++ {
		fmt.Printf("output: %v \n", <-outputs)
	}
	close(outputs)

}
