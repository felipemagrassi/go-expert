package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d got %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	qtWorker := 1000000

	for i := 0; i < qtWorker; i++ {
		go worker(i, data)
	}

	for i := 0; i < qtWorker*10; i++ {
		data <- i
	}

}
