package handlers

type GetAllProductResp struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type AddProductRequset struct {
	//User_Id     int     `json:"userid" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type AddProductResp struct {
	ID      int    `json:"product_id"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
