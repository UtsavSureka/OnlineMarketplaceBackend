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
