package productdb

import (
	DBConnection "Ecomm/pkg/db/DbConnection"
	"Ecomm/pkg/models"
	"errors"
)

func GetAllProducts() ([]models.Product, error) {
	DB, err := DBConnection.DBConnection()
	if err != nil {
		return []models.Product{}, err
	}

	var Products []models.Product

	query := "SELECT id,name,description,price,quantity FROM Products"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, errors.New("Falied to execute the query " + err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product
		rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)
		Products = append(Products, product)
	}
	return Products, nil
}

func AddNewProduct(product models.Product) (int, error) {
	DB, err := DBConnection.DBConnection()

	if err != nil {
		return 0, errors.New(err.Error())
	}

	defer DB.Close()

	query := "INSERT INTO Products (user_id,name,description,price,quantity) VALUES(?,?,?,?,?)"

	result, err := DB.Exec(query, product.User_Id, product.Name, product.Description, product.Price, product.Quantity)

	if err != nil {
		return 0, errors.New("Unable to insert data into database " + err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("Unable to fetch last inserted id" + err.Error())
	}

	return int(id), nil
}

func UpdateProductDetails(product models.Product) error {

	DB, err := DBConnection.DBConnection()

	if err != nil {
		return errors.New("unable to connect to database")
	}

	query := "UPDATE products SET name=?,description=?,price=?,quantity=?,updated_at=CURRENT_TIMESTAMP WHERE id=?"
	_, err = DB.Exec(query, product.Name, product.Description, product.Price, product.Quantity, product.Id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProductById(id int) error {
	DB, err := DBConnection.DBConnection()
	if err != nil {
		return errors.New("unable to connect to database")
	}

	query := "DELETE FROM products WHERE id=?"
	_, err = DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
