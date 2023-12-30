package main

import "time"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- 2
	}()

	for i := 0; i < 6; i++ {
		select {
		case x := <-c1:
			println("Received", x, "from c1")
		case x := <-c2:
			println("Received", x, "from c2")
		case <-time.After(3 * time.Second):
			println("Timeout")
		default:
			println("Default")
		}
	}

}
