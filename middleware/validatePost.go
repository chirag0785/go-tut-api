package middleware

import (
	"github.com/chirag0785/go-tut-api/dto"

	"github.com/gin-gonic/gin"
)
func ValidatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body dto.PostCreateDTO
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{
				"error": "invalid request body",
			})
			c.Abort()
			return
		}

		c.Set("body", &body)
		c.Next()
	}
}