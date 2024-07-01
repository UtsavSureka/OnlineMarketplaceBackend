package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Build creates and configures a new Gin server instance.
func Build() *gin.Engine {
	router := gin.Default()

	// Ping Example
	// @Summary Ping example
	// @Description Do ping
	// @ID ping
	// @Produce  plain
	// @Success 200 {string} string "pong"
	// @Router /ping [get]
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"msg": "connection established",
		})
	})

	// API v1 endpoint
	v1 := router.Group("/api/v1")

	groupAllEndPoints(v1) // Define your API endpoints here

	return router

}
