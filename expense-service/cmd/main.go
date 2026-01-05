package main

import (
	"fmt"
	"net/http"

	"github.com/Aytaditya/splitnest-expense-service/internal/config"
)

func main() {
	config := config.MustLoad()
	fmt.Println("Loaded config:", config)
	// here we will connect to db
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Expense Service is running"))
	})
	fmt.Println("Starting Expense Service on port 8083")
	err := http.ListenAndServe(":8083", router)
	if err != nil {
		panic(err)
	}
}
