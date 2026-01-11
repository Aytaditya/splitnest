package connection

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Exchange must already exist, but safe to redeclare
	err = ch.ExchangeDeclare(
		"expense.events",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	log.Println("Notification service connected to RabbitMQ")

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}, nil
}

func (r *RabbitMQ) Close() {
	r.Channel.Close()
	r.Conn.Close()
}
