package types

type ExpenseCreatedEvent struct {
	ExpenseID int64   `json:"expense_id"`
	GroupID   int64   `json:"group_id"`
	PaidBy    int64   `json:"paid_by"`
	Amount    int64   `json:"amount"`
	Members   []int64 `json:"members"`
}
