package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello!")
	})

	app.Run(":9999")

}
