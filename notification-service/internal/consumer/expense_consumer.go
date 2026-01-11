package consumer

import (
	"encoding/json"
	"log"

	"github.com/Aytaditya/splitnest-notification-service/internal/types"
	amqp "github.com/rabbitmq/amqp091-go"
)

func StartExpenseConsumer(ch *amqp.Channel) error {

	q, err := ch.QueueDeclare(
		"notification.expense",
		true, // durable
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,
		"expense.created",
		"expense.events",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false, // manual ACK
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Println("Notification consumer started, waiting for messages...")

	go func() {
		for msg := range msgs {
			var event types.ExpenseCreatedEvent

			err := json.Unmarshal(msg.Body, &event)
			if err != nil {
				log.Println("invalid message format:", err)
				msg.Nack(false, false) // drop bad message
				continue
			}

			// 🔔 BUSINESS LOGIC (for now just log)
			log.Printf(
				"Notify: Expense %d created in group %d, amount %d by user %d",
				event.ExpenseID,
				event.GroupID,
				event.Amount,
				event.PaidBy,
			)

			// TODO:
			// - Send email
			// - Push notification
			// - Store notification in DB

			msg.Ack(false)
		}
	}()

	return nil
}
