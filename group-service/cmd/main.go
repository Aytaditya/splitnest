package main

import (
	"fmt"
	"net/http"

	"github.com/Aytaditya/splitnest-group-service/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Printf("Config loaded: %+v\n", cfg)
	// we will connect with db
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	fmt.Println("Starting server on :8082")
	err := http.ListenAndServe(":8082", router)
	if err != nil {
		panic(err)
	}
}
