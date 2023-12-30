package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Msg string
	id  uint64
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	go func() {
		var i uint64 = 0
		for {
			atomic.AddUint64(&i, 1)
			time.Sleep(1 * time.Second)
			msg := Message{"Hello", i}
			c1 <- msg
		}
	}()

	go func() {
		var i uint64 = 0
		for {
			atomic.AddUint64(&i, 1)
			time.Sleep(2 * time.Second)
			msg := Message{"World", i}
			c2 <- msg
		}
	}()

	for {
		select {
		case x := <-c1:
			fmt.Printf("Received %s from c1 ID: %d\n", x.Msg, x.id)
		case x := <-c2:
			fmt.Printf("Received %s from c2 ID: %d\n", x.Msg, x.id)
		case <-time.After(3 * time.Second):
			println("Timeout")
			// default:
			// 	println("Default")
		}
	}

}
