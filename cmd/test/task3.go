package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 100; i++ {
		work(ctx)
		// assume, that at i == 5 some error occurs
		if i == 5 {
			fmt.Printf("cancel of rpcCall:\n")
			cancel()
		}
	}

	fmt.Println("num :", runtime.NumGoroutine())

	// server doesn't die. Imagine, it's doing useful work.
	for {
		fmt.Printf("i do some useful work, print num: %d\n", rand.Int())
		time.Sleep(time.Second)
	}
}

func work(ctx context.Context) {
	ch := resCh()

	go func() {
		ch <- rpcCall()
	}()

	select {
	case value := <-ch:
		fmt.Printf("result of rpcCall: %d\n", value)
	case <-ctx.Done():
		return
	}
}

func rpcCall() int {
	fmt.Println("start rpcCall:")
	time.Sleep(time.Second)
	fmt.Println("end rpcCall:")
	return rand.Int()
}

func resCh() chan int {
	ch := make(chan int, 1)

	return ch
}
