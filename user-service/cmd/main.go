package main

import (
	"fmt"
	"net/http"
)

func main() {
	// here we will load config
	// here wwe will connect db
	// here we will start the server
	router := http.NewServeMux()
	// we will add routes here
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User Service is up and running"))
	})
	fmt.Println("Server is running on port 8081")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		panic(err)
	}
}
