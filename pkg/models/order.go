package models

type Order struct {
	Id     int           `json:"id"`
	UserId int           `json:"userId"`
	Items  []Order_Items `json:"items"`
	Total  float64       `json:"total_price"`
	Status string        `json:"status"`
}

type Order_Items struct {
	Id           int     `json:"id"`
	Order_id     int     `json:"order_id"`
	Product_id   int     `json:"product_id"`
	Product_name string  `json:"Product_name"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
}
