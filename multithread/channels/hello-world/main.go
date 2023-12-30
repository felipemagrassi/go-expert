package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello world!"
		ch <- "Hello again!" // This line is not executed because the channel is unbuffered and the main goroutine is not receiving from it
	}()

	fmt.Println(<-ch)
}
