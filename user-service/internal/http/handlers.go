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
		response.WriteResponse(w, http.StatusOK, map[string]string{"message": "User Service is healthy"})
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

func Login(storage *storage.Sqlite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var details types.Login
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
		id, email, token, err1 := storage.LoginUser(details.Username, details.Password)
		if err1 != nil {
			response.WriteResponse(w, http.StatusUnauthorized, map[string]string{"error": err1.Error()})
			return
		}
		response.WriteResponse(w, http.StatusOK, map[string]interface{}{"message": "user logged in successfully", "token": token, "id": fmt.Sprintf("%d", id), "email": email})
	}
}

func GetUserByUsername(storage *storage.Sqlite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var username string
		username = r.PathValue("username")
		id, email, err1 := storage.FindUsername(username)
		if err1 != nil {
			response.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err1.Error()})
			return
		}
		response.WriteResponse(w, http.StatusOK, map[string]string{"message": "user found successfully", "id": fmt.Sprintf("%d", id), "email": email})
	}
}
