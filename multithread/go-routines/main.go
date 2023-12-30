package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go task("A")

	go task("C")

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("B: %d\n", i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	time.Sleep(time.Second * 6)

}
