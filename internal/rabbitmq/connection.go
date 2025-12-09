package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
)

var Conn *amqp091.Connection
var Channel *amqp091.Channel

func Connect() error {
	var err error

	Conn, err = amqp091.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return err
	}

	Channel, err = Conn.Channel()
	if err != nil {
		return err
	}

	// Declare exchange
	return Channel.ExchangeDeclare(
		"fleet.events",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
}
