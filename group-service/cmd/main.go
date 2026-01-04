package main

import (
	"fmt"
	"net/http"

	"github.com/Aytaditya/splitnest-group-service/internal/config"
	handlers "github.com/Aytaditya/splitnest-group-service/internal/http"
	"github.com/Aytaditya/splitnest-group-service/internal/storage"
)

func main() {
	cfg := config.MustLoad()
	fmt.Printf("Config loaded: %+v\n", cfg)

	storage, err1 := storage.ConnectDB(cfg)
	if err1 != nil {
		panic(err1)
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.Healthy())
	router.HandleFunc("POST /create-group", handlers.CreateGroup(storage))
	router.HandleFunc("POST /add-members/{groupId}", handlers.AddMembers(storage))

	fmt.Println("Starting server on", cfg.HttpServer.Address)
	err := http.ListenAndServe(cfg.HttpServer.Address, router)
	if err != nil {
		panic(err)
	}
}
