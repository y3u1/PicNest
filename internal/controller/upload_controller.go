package controller

import (
	"PicNest/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadController struct {
	UploadService *services.UploadService
}

func NewUploadController() *UploadController {
	return &UploadController{UploadService: services.NewUploadService()}
}

func (us *UploadController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "文件上传失败！",
		})
		return
	}
	var filename string
	filename, err = us.UploadService.UploadImage(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "文件上传成功",
		"url":     filename, // 假设文件名未更改
	})
}
