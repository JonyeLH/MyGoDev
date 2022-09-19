package main

import (
	"fmt"
	"time"
)

func Worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, "started jobs ", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "finished jobs ", j)
		result <- 2 * j
	}
}

func main() {
	const numjobs = 5
	jobs := make(chan int, numjobs)
	result := make(chan int, numjobs)

	for i := 0; i < 3; i++ {
		go Worker(i, jobs, result)
	}
	for j := 1; j < 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a < numjobs; a++ {
		res := <-result
		fmt.Println("result = ", res)
	}
}
