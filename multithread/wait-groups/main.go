package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(time.Millisecond * 500)
		wg.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(30)

	go task("A", &waitGroup)
	go task("C", &waitGroup)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("B: %d\n", i)
			time.Sleep(time.Millisecond * 500)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
}
