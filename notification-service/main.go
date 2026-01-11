package main

import (
	"log"

	"github.com/Aytaditya/splitnest-notification-service/internal/connection"
	"github.com/Aytaditya/splitnest-notification-service/internal/consumer"
)

func main() {
	rmq, err := connection.NewRabbitMQ(
		"amqp://admin:admin123@rabbitmq:5672/",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rmq.Close()

	err = consumer.StartExpenseConsumer(rmq.Channel)
	if err != nil {
		log.Fatal(err)
	}

	// BLOCK FOREVER
	select {}
}
