package middleware

import "github.com/gin-gonic/gin"

func Auth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{"message": "Authorization header is required"})
		c.Abort()
	}
	c.Next()
}
