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
		log.Fatal(err)
	}

	ch, _ := conn.Channel()

	q, _ := ch.QueueDeclare(
		"geofence_alerts",
		true,
		false,
		false,
		false,
		nil,
	)

	ch.QueueBind(
		q.Name,
		"geofence.alert",
		"fleet.events",
		false,
		nil,
	)

	msgs, _ := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	log.Println("Worker started...")

	for msg := range msgs {
		log.Println("⚠️ GEOFENCE EVENT:", string(msg.Body))
	}
}
