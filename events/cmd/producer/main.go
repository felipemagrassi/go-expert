package main

import (
	"fmt"
	"log"

	"github.com/felipemagrassi/go-expert/events/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		log.Fatalf("Error opening channel: %s", err)
		panic(err)
	}

	defer ch.Close()

	for i := 0; i < 10000; i++ {
		msg := fmt.Sprintf("Hello world! %d", i)
		rabbitmq.Publish(ch, []byte(msg))
	}
}
