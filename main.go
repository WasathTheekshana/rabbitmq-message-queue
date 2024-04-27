package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Hello rabbit")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("Failed to connect to RabbitMQ: %s", err)
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Successfully connected to RabbitMQ 🐰")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Failed to open a channel: %s", err)
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Test",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Printf("Failed to declare a queue: %s", err)
		panic(err)
	}

	fmt.Println("Queue created: ", q.Name)

	err = ch.Publish(
		"",
		"Test",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte("Hello World It's Wasath 🌍"),
		},
	)

	if err != nil {
		fmt.Printf("Failed to publish a message: %s", err)
		panic(err)
	}

	fmt.Println("Message published 🚀")

	q, err = ch.QueueInspect("Test")
	if err != nil {
		fmt.Printf("Failed to inspect the queue: %s", err)
		panic(err)
	}

	fmt.Println("Message count: ", q.Messages)
}