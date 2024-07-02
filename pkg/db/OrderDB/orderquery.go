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
