package main

import (
	"fmt"
	"net/http"

	"github.com/Aytaditya/splitnest-expense-service/internal/config"
	handlers "github.com/Aytaditya/splitnest-expense-service/internal/http"
	"github.com/Aytaditya/splitnest-expense-service/internal/storage"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println("Loaded config:", cfg)
	storage, err2 := storage.ConnectDB(cfg)
	if err2 != nil {
		panic(err2)
	}
	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.Healthy())
	router.HandleFunc("POST /add-expense/{groupId}", handlers.AddExpense(storage))

	fmt.Println("Starting Expense Service on port 8083")
	err := http.ListenAndServe(cfg.HttpServer.Address, router)
	if err != nil {
		panic(err)
	}
}
