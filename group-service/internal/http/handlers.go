package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	call "github.com/Aytaditya/splitnest-group-service/internal/http/sync-call"
	"github.com/Aytaditya/splitnest-group-service/internal/middleware"
	"github.com/Aytaditya/splitnest-group-service/internal/response"
	"github.com/Aytaditya/splitnest-group-service/internal/storage"
)

func Healthy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteResponse(w, http.StatusOK, map[string]string{"message": "Group Service is healthy"})
	}
}

func CreateGroup(storage *storage.Sqlite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := middleware.ValidateToken(r.Header.Get("Authorization"))
		if err != nil {
			response.WriteResponse(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}
		fmt.Println("User ID from token:", id)

		var details struct {
			Name string `json:"name"`
		}
		err2 := json.NewDecoder(r.Body).Decode(&details)
		if errors.Is(err2, io.EOF) {
			response.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
			return
		}
		if err2 != nil {
			response.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
			return
		}
		groupId, err3 := storage.CreateGroup(details.Name, id)
		if err3 != nil {
			response.WriteResponse(w, http.StatusInternalServerError, err3.Error())
			return
		}

		response.WriteResponse(w, http.StatusOK, map[string]string{"message": "Group created successfully", "group_id": fmt.Sprintf("%d", groupId)})
	}
}

// in this function first validate that the requester that he is owner of group and then by username get user id which he is adding to the group and also check if user already exists in the group
func AddMembers(storage *storage.Sqlite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var details struct {
			Username string `json:"username"`
		}
		err := json.NewDecoder(r.Body).Decode(&details)
		if errors.Is(err, io.EOF) {
			response.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
			return
		}
		if err != nil {
			response.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
			return
		}
		id, err2 := call.GetUserByUsername(details.Username)
		if err2 != nil {
			response.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err2.Error()})
			return
		}
		fmt.Println("Fetched user ID:", id)

		groupdId := r.PathValue("groupId")
		groupId_int, err3 := strconv.ParseInt(groupdId, 10, 64)
		if err3 != nil {
			response.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid group ID"})
			return
		}
		fmt.Println("Group ID from path:", groupId_int)

		// jwt token
		requesterId, err4 := middleware.ValidateToken(r.Header.Get("Authorization"))
		if err4 != nil {
			response.WriteResponse(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}
		err5 := storage.AddMembers(groupId_int, requesterId, id)
		if err5 != nil {
			response.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err5.Error()})
			return
		}

		response.WriteResponse(w, http.StatusOK, map[string]string{"message": "Member added successfully"})
	}
}
