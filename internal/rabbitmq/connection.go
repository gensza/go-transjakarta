package rabbitmq

import (
	"os"

	"github.com/rabbitmq/amqp091-go"
)

var Conn *amqp091.Connection
var Channel *amqp091.Channel

func Connect() error {
	var err error

	Conn, err = amqp091.Dial(os.Getenv("RABBITMQ_HOST"))
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
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
}
