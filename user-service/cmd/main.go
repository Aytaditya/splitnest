package main

import (
	"fmt"
	"net/http"

	"github.com/Aytaditya/splitnest-user-service/internal/config"
	handlers "github.com/Aytaditya/splitnest-user-service/internal/http"
	"github.com/Aytaditya/splitnest-user-service/internal/storage"
)

func main() {
	// here we will load config
	cfg, err1 := config.LoadConfig()
	if err1 != nil {
		panic(err1)
	}
	//fmt.Println("Loaded config:", cfg)

	// here wwe will connect db
	storage, err2 := storage.ConnectDB(cfg)
	if err2 != nil {
		panic(err2)
	}
	// here we will start the server
	router := http.NewServeMux()
	// we will add routes here
	router.HandleFunc("GET /", handlers.Healthy())
	router.HandleFunc("POST /register", handlers.Signup(storage))

	fmt.Println("Server is running on port 8081")
	err := http.ListenAndServe(cfg.HttpServer.Address, router)
	if err != nil {
		panic(err)
	}
}
