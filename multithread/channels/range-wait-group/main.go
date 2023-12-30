package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(10)

	go publish(ch, &wg)
	go subscribe(ch, &wg)

	wg.Wait()
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
}

func subscribe(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Printf("Received %d\n", i)
		wg.Done()
	}

}
