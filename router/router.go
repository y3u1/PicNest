package router

import (
	"PicNext/router/api"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/upload", api.UploadImage)
		//v1.GET("/list")
	}
	return r
}
