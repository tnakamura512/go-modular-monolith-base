package mcon

type SendMessageRequest struct {
	Token   string `json:"token"`
	Message string `json:"message"`
	SendIds []int  `json:"send_user_ids"`
}
