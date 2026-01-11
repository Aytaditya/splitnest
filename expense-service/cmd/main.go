package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aytaditya/splitnest-expense-service/internal/config"
	"github.com/Aytaditya/splitnest-expense-service/internal/events"
	handlers "github.com/Aytaditya/splitnest-expense-service/internal/http"
	"github.com/Aytaditya/splitnest-expense-service/internal/storage"
)

func main() {
	cfg := config.MustLoad()
	publisher, err3 := events.NewPublisher("amqp://admin:admin123@rabbitmq:5672/")
	if err3 != nil {
		log.Fatal(err3)
	}

	defer publisher.Close()
	fmt.Println("Loaded config:", cfg)
	storage, err2 := storage.ConnectDB(cfg)
	if err2 != nil {
		panic(err2)
	}
	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.Healthy())
	router.HandleFunc("POST /add-expense/{groupId}", handlers.AddExpense(storage, publisher))
	router.HandleFunc("GET /get-expense/{groupId}", handlers.GetExpenses(storage))

	fmt.Println("Starting Expense Service on port 8083")
	err := http.ListenAndServe(cfg.HttpServer.Address, router)
	if err != nil {
		panic(err)
	}
}
