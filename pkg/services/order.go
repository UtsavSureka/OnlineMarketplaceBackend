package services

import (
	DBConnection "Ecomm/pkg/db/DbConnection"
	orderdb "Ecomm/pkg/db/OrderDB"
	productdb "Ecomm/pkg/db/ProductDB"
	"Ecomm/pkg/models"
	"fmt"
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
	DB, err := DBConnection.DBConnection()
	if err != nil {
		return models.Order{}, err
	}

	//Map to store orderid and bool to find if the order with order id is present or not
	isPresent := make(map[int]bool)

	rows, err := DB.Query("SELECT id FROM orders")
	if err != nil {
		return models.Order{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var orderId int
		rows.Scan(&orderId)
		isPresent[orderId] = true
	}

	if !isPresent[id] {
		return models.Order{}, fmt.Errorf("order with order id : %d is not present", id)
	}

	orderDetail, err := orderdb.GetOrderByOrderId(id)
	if err != nil {
		return models.Order{}, err
	}

	return orderDetail, nil
}
