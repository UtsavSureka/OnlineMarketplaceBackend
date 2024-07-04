package orderdb

import (
	DBConnection "Ecomm/pkg/db/DbConnection"
	"Ecomm/pkg/models"
	"errors"
)

func CreataOrder(order models.Order) (int, error) {

	DB, err := DBConnection.DBConnection()
	if err != nil {
		return 0, errors.New("error connecting to database")
	}

	defer DB.Close()

	orderQuery := "INSERT INTO ORDERS (user_id,total_amount) VALUES (?,?)"

	result, err := DB.Exec(orderQuery, order.UserId, order.Total)
	if err != nil {
		return 0, err
	}

	orderId, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("unable to fetch last inserted id from db")
	}

	for i := range order.Items {
		order.Items[i].Order_id = int(orderId)
	}

	orderItemsQuery := "INSERT INTO orderItems (order_id,product_id,product_name,quantity,price) VALUES (?,?,?,?,?)"

	for i := range order.Items {
		_, err = DB.Exec(orderItemsQuery, order.Items[i].Order_id, order.Items[i].Product_id, order.Items[i].Product_name, order.Items[i].Quantity, order.Items[i].Price)
		if err != nil {
			_, err = DB.Exec("DELETE FROM orders WHERE id = ?", int(orderId))
			return 0, err
		}

	}

	//As order is created successfully we need to change the current quantity in product table

	return int(orderId), nil
}

func GetOrderByOrderId(id int) (models.Order, error) {

	DB, err := DBConnection.DBConnection()
	if err != nil {
		return models.Order{}, err
	}
	// creating a variable orderDetail to store all the data of order
	var orderDetail models.Order

	orderRow := DB.QueryRow("SELECT id,user_id,total_amount,status FROM orders WHERE id=?", id)
	err = orderRow.Scan(&orderDetail.Id, &orderDetail.UserId, &orderDetail.Total, &orderDetail.Status)
	if err != nil {
		return models.Order{}, err
	}

	orderItems, err := DB.Query("SELECT id,order_id,product_id,product_name,quantity,price FROM orderitems WHERE order_id=?", id)
	if err != nil {
		return models.Order{}, err
	}

	defer orderItems.Close()

	for orderItems.Next() {
		var orderItem models.Order_Items
		orderItems.Scan(&orderItem.Id, &orderItem.Order_id, &orderItem.Product_id, &orderItem.Product_name, &orderItem.Quantity, &orderItem.Price)
		orderDetail.Items = append(orderDetail.Items, orderItem)
	}

	return orderDetail, nil

}

func GetAllOrdersByUserId(id int) ([]models.Order, error) {
	DB, err := DBConnection.DBConnection()
	if err != nil {
		return []models.Order{}, err
	}

	var orders []models.Order
	orderRows, err := DB.Query("SELECT id,user_id,total_amount,status FROM orders")
	if err != nil {
		return []models.Order{}, err
	}

	orderItemsQuery := "SELECT id,order_id,product_id,product_name,quantity,price FROM OrderItems WHERE order_id=?"

	for orderRows.Next() {
		var order models.Order
		orderRows.Scan(&order.Id, &order.UserId, &order.Total, &order.Status)

		var orderItems []models.Order_Items
		orderItemRows, err := DB.Query(orderItemsQuery, order.Id)
		if err != nil {
			return []models.Order{}, err
		}
		for orderItemRows.Next() {
			var orderItem models.Order_Items
			orderItemRows.Scan(&orderItem.Id, &orderItem.Order_id, &orderItem.Product_id, &orderItem.Product_name, &orderItem.Quantity, &orderItem.Price)
			orderItems = append(orderItems, orderItem)
		}
		order.Items = append(order.Items, orderItems...)
		orders = append(orders, order)
	}

	return orders, nil

}

func CancelOrderByOrderId(id int) error {

	DB, err := DBConnection.DBConnection()
	if err != nil {
		return err
	}

	defer DB.Close()
	_, err = DB.Exec("UPDATE orders SET status=? WHERE id=?", "cancelled", id)
	if err != nil {
		return err
	}

	// Write a function to update the product quantity as order is cancelled :

	orderItems, err := DB.Query("SELECT product_id,quantity FROM orderitems where order_id=?", id)
	if err != nil {
		return err
	}

	type OrderQuantity struct {
		product_id int
		quantity   int
	}

	var orderQuantity []OrderQuantity

	for orderItems.Next() {
		var orders OrderQuantity
		orderItems.Scan(&orders.product_id, &orders.quantity)
		orderQuantity = append(orderQuantity, orders)
	}

	for i := range orderQuantity {
		var productQuantity int
		row := DB.QueryRow("SELECT quantity FROM products WHERE id=?", orderQuantity[i].product_id)
		row.Scan(&productQuantity)

		_, err := DB.Exec("UPDATE products SET quantity=? WHERE id=?", productQuantity-orderQuantity[i].quantity, orderQuantity[i].product_id)
		if err != nil {
			return err
		}
	}
	return nil
}
