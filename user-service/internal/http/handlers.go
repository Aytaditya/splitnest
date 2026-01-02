package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Aytaditya/splitnest-user-service/internal/response"
	"github.com/Aytaditya/splitnest-user-service/internal/storage"
	"github.com/Aytaditya/splitnest-user-service/internal/types"
)

func Healthy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User Service is up and running"))
	}
}

func Signup(storage *storage.Sqlite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var details types.Signup
		err := json.NewDecoder(r.Body).Decode(&details)
		if errors.Is(err, io.EOF) {
			response.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "empty request body"})
			return
		}
		if err != nil {
			response.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
			return
		}
		fmt.Println(details)
		id, token, err1 := storage.RegisterUser(details.Username, details.Email, details.Password)
		if err1 != nil {
			response.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err1.Error()})
			return
		}
		response.WriteResponse(w, http.StatusCreated, map[string]interface{}{"message": "user registered successfully", "user_id": id, "token": token})

	}
}
