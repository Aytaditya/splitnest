package call

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Aytaditya/splitnest-expense-service/internal/types"
)

func GetMembers(groupId string) ([]int64, error) {
	url := "http://group-service:8082/group-members/" + groupId
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err2 := client.Do(req)
	if err2 != nil {
		return nil, err2
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("user not found due to unknown error")
	}

	var members []types.GetMemberResponse
	err3 := json.NewDecoder(res.Body).Decode(&members)
	if err3 != nil {
		return nil, err3
	}

	userIDs := make([]int64, 0, len(members))
	for _, m := range members {
		userIDs = append(userIDs, m.UserId)
	}

	return userIDs, nil
}
