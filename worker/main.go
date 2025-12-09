package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
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
