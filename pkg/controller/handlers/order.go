package handlers

import (
	"Ecomm/pkg/models"
	"Ecomm/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateNewOrderHandler(ctx *gin.Context) {
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

func GetOrderByIdHandler(ctx *gin.Context) {
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

func GetOrderHistoryHandler(ctx *gin.Context) {
	user_id, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "Unable to get user id",
		})
		return
	}
	id := user_id.(int)

	orders, err := services.GetAllOrdersByUserId(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

func CancelOrderHandler(ctx *gin.Context) {
	orderId := ctx.Query("id")
	if orderId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}
	id, err := strconv.Atoi(orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid order id",
		})
		return
	}
	err = services.CancelAnOrder(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order Cancelled",
	})
}
