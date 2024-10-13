package controller

type AddOrderRequest struct {
	ProductId int `json:"product_id"`
}

type PayOrderRequest struct {
	OrderId int `json:"id"`
}
