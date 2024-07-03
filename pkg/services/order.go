package services

import (
	DBConnection "Ecomm/pkg/db/DbConnection"
	orderdb "Ecomm/pkg/db/OrderDB"
	productdb "Ecomm/pkg/db/ProductDB"
	"Ecomm/pkg/models"
	"errors"
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
	productQuanttityMapping := make(map[int]int)

	for _, product := range products {
		productIdMapping[product.Id] = product.Price
		productNameMapping[product.Id] = product.Name
		productQuanttityMapping[product.Id] = product.Quantity
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

	//As order is created successfully we need to change the current quantity in product table
	for i := range order.Items {
		currentQuantity := productQuanttityMapping[order.Items[i].Product_id] - order.Items[i].Quantity
		DB, err := DBConnection.DBConnection()
		if err != nil {
			return 0, errors.New("unable to make DB connection for updating product")
		}
		_, err = DB.Exec("UPDATE products SET quantity=?,updated_at=CURRENT_TIMESTAMP WHERE id = ?", currentQuantity, order.Items[i].Product_id)
		if err != nil {
			return 0, errors.New("unable to update product quantity")
		}
		defer DB.Close()
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

func GetAllOrdersByUserId(id int) ([]models.Order, error) {
	DB, err := DBConnection.DBConnection()
	if err != nil {
		return []models.Order{}, errors.New("unable to make connectio to Database")
	}

	defer DB.Close()
	//Check if order with the user id given exists or not
	var isPresent bool
	rows, err := DB.Query("SELECT user_id FROM orders WHERE user_id=?", id)
	if err != nil {
		return []models.Order{}, err
	}
	isPresent = rows.Next()

	if !isPresent {
		return []models.Order{}, fmt.Errorf("no order present with order id %d", id)
	}

	orders, err := orderdb.GetAllOrdersByUserId(id)
	if err != nil {
		return []models.Order{}, err
	}

	return orders, nil

}

func CancelAnOrder(id int) error {
	DB, err := DBConnection.DBConnection()
	if err != nil {
		return err
	}

	defer DB.Close()

	rows, err := DB.Query("SELECT id From orders WHERE id=?", id)
	if err != nil {
		return err
	}

	if !rows.Next() {
		return fmt.Errorf("order with order id %d is not present", id)
	}

	err = orderdb.CancelOrderByOrderId(id)
	if err != nil {
		return err
	}

	return nil

}
