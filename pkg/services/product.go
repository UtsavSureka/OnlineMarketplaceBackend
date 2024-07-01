package services

import (
	productdb "Ecomm/pkg/db/ProductDB"
	"Ecomm/pkg/models"
	"errors"
	"fmt"
)

func GetAllProduct() ([]models.Product, error) {

	Products, err := productdb.GetAllProducts()
	if err != nil {
		return []models.Product{}, errors.New("unable to fetch product details from database")
	}
	return Products, nil

}

func AddNewProduct(product models.Product) (int, error) {

	id, err := productdb.AddNewProduct(product)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetProductById(id int) (models.Product, error) {
	Products, err := productdb.GetAllProducts()
	if err != nil {
		return models.Product{}, errors.New("unable to fetch product details from database")
	}

	for _, product := range Products {
		if product.Id == id {
			return product, nil
		}
	}

	return models.Product{}, fmt.Errorf("product with id %d not found", id)
}

func UpdateProductById(newProduct models.Product) error {
	Products, err := GetAllProduct()
	if err != nil {
		return err
	}

	var isFound bool
	for _, product := range Products {
		if product.Id == newProduct.Id {
			isFound = true
			break
		}
	}
	if !isFound {
		return fmt.Errorf("product with id %d not found", newProduct.Id)
	}

	err = productdb.UpdateProductDetails(newProduct)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProductById(id int) error {
	Products, err := GetAllProduct()
	if err != nil {
		return err
	}

	var isFound bool
	for _, product := range Products {
		if product.Id == id {
			isFound = true
			break
		}
	}
	if !isFound {
		return fmt.Errorf("product with id %d not found", id)
	}

	err = productdb.DeleteProductById(id)
	if err != nil {
		return err
	}
	return nil
}
