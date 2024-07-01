package handlers

import (
	"Ecomm/pkg/models"
	"Ecomm/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateNewOrder(ctx *gin.Context) {
	user_id, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch user id",
		})
		return
	}

	var req models.Order
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Payload",
		})
		return
	}

	req.UserId = user_id.(int)

	id, err := services.CreateOrder(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Order placed successfully",
		"Order Id": id,
	})

}

func GetOrderById(ctx *gin.Context) {
	id_string := ctx.Query("id")
	if id_string == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}
	id, err := strconv.Atoi(id_string)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid order id",
		})
		return
	}

	resp, err := services.GetOrderDetailsById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)

}
