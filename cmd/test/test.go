package main

import (
	"fmt"
)

func main3() {

	job := make(chan int, 4)

	out := task(job)

	for i := 0; i < 12; i++ {
		job <- i
	}
	close(job)

	for v := range out {
		fmt.Println(v)
	}

}

func task(job chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range job {
			out <- v
		}
		close(out)
	}()

	return out
}
