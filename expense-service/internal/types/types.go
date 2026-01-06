package types

type GetMemberResponse struct {
	UserId int64 `json:"user_id"`
}

type BalanceUser struct {
	UserId int64 `json:"user_id"`
	Amount int64 `json:"amount"`
}
