package main

import (
	"fmt"
	"time"
)

//goroutine pool

func work(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d over job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {

	job := make(chan int, 100)
	result := make(chan int, 100)

	//goroutine pool
	for i := 1; i <= 10; i++ {
		go work(i, job, result)
	}

	//20个任务
	for j := 1; j <= 20; j++ {
		job <- j
	}
	close(job)

	//输出

	for y := 1; y <= 19; y++ {
		<-result
	}
}
