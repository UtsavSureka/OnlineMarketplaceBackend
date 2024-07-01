package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Get the token from request header

		tokenString := ctx.GetHeader("Authorization")

		//Check if token is missing

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized access",
			})
			ctx.Abort()
			return
		}

		claims, err := ValidateJwtToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}
		ctx.Set("id", claims.Id)
		ctx.Next()

	}

}

func IsAdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Get the token from request header

		tokenString := ctx.GetHeader("Authorization")

		//Check if token is missing

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized access",
			})
			ctx.Abort()
			return
		}

		claims, err := ValidateJwtToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		isAdmin := claims.IsAdmin

		if !isAdmin {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "This request is only valid for admin type user"})
			ctx.Abort()
			return
		}
		ctx.Set("id", claims.Id)
		ctx.Next()

	}

}
