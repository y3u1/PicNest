package middleware

import (
	"PicNest/internal/services"

	"github.com/gin-gonic/gin"
)

func Auth(as *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"message": "Authorization header is required"})
			c.Abort()
		}
		err := as.Authenticate(token)
		if err != nil {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
		} else {
			c.Set("user", token) // Store the user info in the context
			// Proceed to the next handler
		}
		c.Next()
	}
}
