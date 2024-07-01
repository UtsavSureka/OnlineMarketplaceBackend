package services

import (
	orderdb "Ecomm/pkg/db/OrderDB"
	productdb "Ecomm/pkg/db/ProductDB"
	"Ecomm/pkg/models"
)

func CreateOrder(order models.Order) (int, error) {

	products, err := productdb.GetAllProducts()
	if err != nil {
		return 0, err
	}

	//Make a map of id and respective price
	productIdMapping := make(map[int]float64)
	productNameMapping := make(map[int]string)
	for _, product := range products {
		productIdMapping[product.Id] = product.Price
		productNameMapping[product.Id] = product.Name
	}

	//Calculate the total value of order and map order price to product price
	var total float64

	for i := range order.Items {
		order.Items[i].Price = productIdMapping[order.Items[i].Product_id]
		order.Items[i].Product_name = productNameMapping[order.Items[i].Product_id]
		total += order.Items[i].Price * float64(order.Items[i].Quantity)
	}

	order.Total = total

	order_id, err := orderdb.CreataOrder(order)
	if err != nil {
		return 0, err
	}
	return order_id, nil
}

func GetOrderDetailsById(id int) (models.Order, error) {

	return models.Order{}, nil
}
