package main

import (
	_ "Ecomm/docs" // Import the generated docs
	"Ecomm/server"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Documenting API (Online Marketplace Apis)
// @Version 2.0
// @Description API for Online Marketplace having several functionality to handle user,product and order details.
// @contact.name Utsav Sureka
// @contact.url utsav.sureka@stl.tech
// @contact.email ramsureka007@gmail.com
// @Host localhost:5000
// @BasePath /api/v1
func main() {

	// Initialize the Gin server
	server := server.Build()
	// Serve the Swagger documentation
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Run the server
	err := server.Run(":5000")
	if err != nil {
		log.Fatal(err)
	}
}
