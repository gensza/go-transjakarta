package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main() {

	godotenv.Load()

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_HOST"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"geofence_alerts", // queue name
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)

	if err != nil {
		panic(err)
	}

	log.Println("Waiting for messages...")
	for msg := range msgs {
		log.Printf("Received: %s", msg.Body)
	}
}
