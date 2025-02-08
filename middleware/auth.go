package middleware

import (
	"github.com/gin-gonic/gin"

	"my_diary/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		userId := utils.ParseToken(authHeader)
		c.Set("userId", userId)
		c.Next() 
	}
}