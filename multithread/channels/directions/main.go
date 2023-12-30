package main

func hello(ch chan<- string) {
	ch <- "Hello world!"
	// <-ch Invalid operation: <-ch (receive from send-only type chan<- string)
}

func world(ch <-chan string) {
	// ch <- "Hello world!" Invalid operation: ch <- "Hello world!" (send to receive-only type <-chan string)
	println(<-ch)
}

func main() {
	ch := make(chan string)

	go hello(ch)

	world(ch)
}
