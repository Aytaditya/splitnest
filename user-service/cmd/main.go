package main

import (
	"fmt"
	"net/http"

	"github.com/Aytaditya/splitnest-user-service/internal/config"
)

func main() {
	// here we will load config
	cfg, err1 := config.LoadConfig()
	if err1 != nil {
		panic(err1)
	}
	//fmt.Println("Loaded config:", cfg)

	// here wwe will connect db
	// here we will start the server
	router := http.NewServeMux()
	// we will add routes here
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User Service is up and running"))
	})
	fmt.Println("Server is running on port 8081")
	err := http.ListenAndServe(cfg.HttpServer.Address, router)
	if err != nil {
		panic(err)
	}
}
