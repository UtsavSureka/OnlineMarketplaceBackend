package handlers

import (
	"Ecomm/pkg/models"
	"Ecomm/pkg/services"
	"net/http"

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
