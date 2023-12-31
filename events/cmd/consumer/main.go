package main

import (
	"log"

	"github.com/felipemagrassi/go-expert/events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		log.Fatalf("Error opening channel: %s", err)
		panic(err)
	}

	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consumer(ch, msgs)

	for msg := range msgs {
		msg.Ack(false)
	}

}
