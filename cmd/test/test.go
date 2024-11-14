package main

import (
	"encoding/json"
	"fmt"
)

type Name struct {
	Mem int `json:"mem"`
}

func main() {
	var dd int
	myVal := Name{}
	bytes := `{"mem":"1ss"}`
	err := json.Unmarshal([]byte(bytes), &myVal)
	if err != nil {
	}

	fmt.Println(dd)
	//job := make(chan int, 4)
	//
	//out := task(job)
	//
	//for i := 0; i < 12; i++ {
	//	job <- i
	//}
	//close(job)
	//
	//for v := range out {
	//	fmt.Println(v)
	//}

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
