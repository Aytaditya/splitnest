package call

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Aytaditya/splitnest-group-service/internal/types"
)

// we will replace this with grpc
func GetUserByUsername(username string) (int64, error) {
	url := "http://user-service:8081/find-user/" + username
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil) // method url and body
	if err != nil {
		return 0, err
	}
	res, err2 := client.Do(req) // do sends http request and returns http response
	if err2 != nil {
		return 0, err2
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return 0, errors.New("user not found")
	}

	if res.StatusCode != http.StatusOK {
		return 0, errors.New("user not found due to unknown error")
	}

	var details types.UserIdResponse
	err3 := json.NewDecoder(res.Body).Decode(&details)
	if err3 != nil {
		return 0, err3
	}
	userId, err4 := strconv.ParseInt(details.Id, 10, 64)
	if err4 != nil {
		return 0, err4
	}
	return userId, nil
}
