package middleware

import (
	"os"
	"strings"

	"github.com/chirag0785/go-tut-api/utils"
	"github.com/gin-gonic/gin"
)


func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authentication logic will go here
		//extract the token from the request header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"error": "missing authorization header",
			})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader,"Bearer ")
		tokenString = strings.TrimSpace(tokenString)
		//validate the token
		claims, err := utils.ValidateJWTToken(tokenString, os.Getenv("JWT_SECRET"))
		if err != nil {
			c.JSON(401, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}
		//if valid, set the user in the context
		c.Set("user_id", uint(claims["user_id"].(float64)))
		//if not valid, return 401 Unauthorized
		c.Next()
	}
}