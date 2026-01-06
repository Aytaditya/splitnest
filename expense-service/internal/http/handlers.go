package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	call "github.com/Aytaditya/splitnest-expense-service/internal/http/sync-call"
	"github.com/Aytaditya/splitnest-expense-service/internal/middleware"
	"github.com/Aytaditya/splitnest-expense-service/internal/response"
	"github.com/Aytaditya/splitnest-expense-service/internal/storage"
)

func Healthy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteResponse(w, http.StatusOK, map[string]string{"status": "Expense Service is healthy"})
	}
}

func AddExpense(storage *storage.Sqlite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// first we need to verify the jwt and get payerId and also check if he exists in the group
		userId, err := middleware.ValidateToken(r.Header.Get("Authorization"))
		if err != nil {
			response.WriteResponse(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}
		fmt.Println("Authenticated user ID:", userId)

		// reading groupid from url
		GroupId := r.PathValue("groupId")
		fmt.Println("Group ID from URL:", GroupId)

		// reading json body for amount
		var details struct {
			Amount int64 `json:"amount"`
		}
		err3 := json.NewDecoder(r.Body).Decode(&details)
		if err3 != nil {
			response.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
			return
		}

		userIds, err4 := call.GetMembers(GroupId)
		if err4 != nil {
			response.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err4.Error()})
			return
		}
		fmt.Println("Group Members User IDs:", userIds)

		// now we can check requester is in userIds
		isMember := false
		for _, id := range userIds {
			if id == userId {
				isMember = true
				break
			}
		}
		if !isMember {
			response.WriteResponse(w, http.StatusForbidden, map[string]string{"error": "User not in group"})
			return
		}

		// now we will calculate the equal split
		splitAmount := details.Amount / int64(len(userIds))
		fmt.Println("Each member should pay:", splitAmount)

		response.WriteResponse(w, http.StatusOK, map[string]string{"status": "Add Expense endpoint"})
	}
}
