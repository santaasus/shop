package controller

type NotificationRequest struct {
	OrderId   int    `json:"order_id"`
	UserEmail string `json:"user_email"`
}
