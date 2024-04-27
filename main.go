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
}
