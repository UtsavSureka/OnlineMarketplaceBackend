package models

type Product struct {
	Id          int     `json:"id"`
	User_Id     int     `json:"userid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
