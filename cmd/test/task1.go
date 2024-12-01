package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type job struct {
	value int64
	state State
}

type State int

const (
	InitialState State = iota
	FirstStage
	SecondStage
	FinishedStage
)

//1.9219057s

func FirstProcessing(in <-chan job) <-chan job {
	out := make(chan job)
	go func() {
		for job := range in {
			job.value = int64(float64(job.value) * math.Pi)
			job.state = FirstStage
			out <- job
		}
		close(out)
	}()

	return out
}

func SecondProcessing(in <-chan job) <-chan job {
	out := make(chan job)
	go func() {
		for job := range in {
			job.value = int64(float64(job.value) * math.E)
			job.state = SecondStage
			out <- job
		}
		close(out)
	}()
	return out
}

func LastProcessing(in <-chan job) <-chan job {
	out := make(chan job)

	go func() {
		for job := range in {
			job.value = int64(float64(job.value) / float64(rand.Intn(10)))
			job.state = FinishedStage
			out <- job
		}
	}()

	return out
}

func main() {
	length := 50_000_000
	start := time.Now()

	jobs := make([]job, length)
	in := make(chan job, length)
	go func() {
		for i := 0; i < length; i++ {
			jobs[i].value = int64(i)
			in <- jobs[i]
		}
		close(in) //178.7847ms

	}()

	out := LastProcessing(
		SecondProcessing(
			FirstProcessing(in),
		),
	)
	finished := time.Since(start)

	fmt.Println(finished)

	fmt.Println(out)
}
