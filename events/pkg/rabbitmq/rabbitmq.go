package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func Consumer(ch *amqp.Channel, out chan<- amqp.Delivery) error {
	msgs, err := ch.Consume(
		"queue",    // queue
		"consumer", // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil)        // args

	if err != nil {
		return err
	}

	for msg := range msgs {
		log.Printf("Received message with message: %s", msg.Body)
		out <- msg
	}

	return nil
}

func Publish(ch *amqp.Channel, body []byte) error {
	err := ch.Publish(
		"amq.direct", // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate

		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	return err
}
