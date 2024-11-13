package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// 1 add limiter
func main2() {
	count := 1000
	workerCount := 10
	jobs := make(chan int, count)
	results := make(chan int, 100)

	//ch := make(chan int, count)

	for i := 0; i < workerCount; i++ {
		go taskWorker(jobs, results)
	}

	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		jobs <- RPCCall()
	}

	close(jobs)
	wg.Add(1)
	go func() {
		for value := range results {
			fmt.Println(value)
		}
		wg.Done()
	}()
	wg.Wait()

}

func taskWorker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j
	}
	close(results)
}

func RPCCall() int {
	return rand.Int()
}
