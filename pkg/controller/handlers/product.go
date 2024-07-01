package handlers

import (
	"Ecomm/pkg/models"
	"Ecomm/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProductsHandler(ctx *gin.Context) {
	var resp []GetAllProductResp

	products, err := services.GetAllProduct()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	for _, product := range products {
		res := GetAllProductResp{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
		}
		resp = append(resp, res)
	}

	ctx.JSON(http.StatusOK, resp)
}

func AddNewProductHandler(ctx *gin.Context) {
	var req AddProductRequset
	var resp AddProductResp

	err := ctx.BindJSON(&req)
	if err != nil {
		resp.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	id, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve user ID"})
		return
	}

	newProduct := models.Product{
		User_Id:     id.(int),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	lastInsertedId, err := services.AddNewProduct(newProduct)

	if err != nil {
		resp.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.ID = lastInsertedId
	resp.Message = "Product added successfully"

	ctx.JSON(http.StatusOK, resp)

}
func GetProductsHandler(ctx *gin.Context) {
	var resp GetAllProductResp

	idString := ctx.Query("id")

	if idString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid User Id",
		})
		return
	}

	product, err := services.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp = GetAllProductResp{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}

	ctx.JSON(http.StatusOK, resp)
}

func UpdateProductHandler(ctx *gin.Context) {
	var req GetAllProductResp
	var newProduct models.Product

	idString := ctx.Query("id")

	if idString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid User Id",
		})
		return
	}
	err = ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect Payload - Bad request",
		})
		return
	}
	newProduct = models.Product{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	err = services.UpdateProductById(newProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product details updated Successfully",
	})
}

func DeleteAProductHandler(ctx *gin.Context) {
	idString := ctx.Query("id")

	if idString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid User Id",
		})
		return
	}
	err = services.DeleteProductById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Product deleted Successfully",
		"id":      id,
	})
}
