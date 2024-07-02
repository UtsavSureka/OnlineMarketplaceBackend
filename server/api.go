package server

import (
	"Ecomm/pkg/controller/handlers"
	"Ecomm/pkg/utils"

	"github.com/gin-gonic/gin"
)

func groupAllEndPoints(v1 *gin.RouterGroup) {
	user := v1.Group("/user")
	{
		user.POST("register", handlers.UserRegistrationHandler)
		user.POST("login", handlers.UserLoginHandler)
		user.GET("Profile", utils.LoginMiddleware(), handlers.GetUserHandler)
		user.PUT("Profile", utils.LoginMiddleware(), handlers.UpdateUserDetails)
	}

	product := v1.Group("/product")
	{
		product.GET("Products", utils.LoginMiddleware(), handlers.GetAllProductsHandler)
		product.POST("product", utils.IsAdminMiddleware(), handlers.AddNewProductHandler)
		product.GET("product", utils.LoginMiddleware(), handlers.GetProductsHandler)
		product.PUT("product", utils.IsAdminMiddleware(), handlers.UpdateProductHandler)
		product.DELETE("product", utils.IsAdminMiddleware(), handlers.DeleteAProductHandler)
	}

	order := v1.Group("/order")
	{
		order.POST("create", utils.LoginMiddleware(), handlers.CreateNewOrder)
		order.GET("OrderDetail", utils.LoginMiddleware(), handlers.GetOrderById)
	}
}
