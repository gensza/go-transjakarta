package rabbitmq

import "github.com/rabbitmq/amqp091-go"

// Publish Event Geofence
func Publish(event []byte) error {
	return Channel.Publish(
		"fleet.events",
		"geofence.alert",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        event,
		},
	)
}

func Close() {
	if Channel != nil {
		Channel.Close()
	}
	if Conn != nil {
		Conn.Close()
	}
}
