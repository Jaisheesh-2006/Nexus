package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Connect connects to RabbitMQ and returns the connection and a channel.
func Connect(url string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	// Declare the exchange
	err = ch.ExchangeDeclare(
		"orders_exchange", // name
		"fanout",          // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, nil, err
	}

	log.Println("Successfully connected to RabbitMQ")
	return conn, ch, nil
}
