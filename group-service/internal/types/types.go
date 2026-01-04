package types

type UserIdResponse struct {
	Email   string `json:"email"`
	Id      string `json:"id"`
	Message string `json:"message"`
}

type UserGroup struct {
	GroupId int64  `json:"group_id"`
	Name    string `json:"group_name"`
}

type GroupMember struct {
	UserId int64 `json:"user_id"`
}
