package events

import (
	"encoding/json"
	"log"

	"github.com/Aytaditya/splitnest-expense-service/internal/types"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewPublisher(rabbitURL string) (*Publisher, error) {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Declare exchange once
	err = ch.ExchangeDeclare(
		"expense.events",
		"topic",
		true, // durable
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	log.Println("RabbitMQ publisher connected")

	return &Publisher{
		conn:    conn,
		channel: ch,
	}, nil
}

func (p *Publisher) PublishExpenseCreated(event types.ExpenseCreatedEvent) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.channel.Publish(
		"expense.events",
		"expense.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (p *Publisher) Close() {
	p.channel.Close()
	p.conn.Close()
}
