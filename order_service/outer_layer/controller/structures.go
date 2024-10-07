package controller

type AddOrderRequest struct {
	UserId    int `json:"user_id"`
	ProductId int `json:"product_id"`
}

type PayOrderRequest struct {
	OrderId int `json:"id"`
}
