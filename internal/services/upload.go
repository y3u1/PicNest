package services

import (
	"PicNest/internal/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type UploadService struct {
}

func (up *UploadService) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "文件上传失败！",
		})
	}
	filename := utils.RandName()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "不允许的文件类型",
		})
	}
	savepath := filepath.Join("uploads", filename)
	if err := c.SaveUploadedFile(file, savepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "文件上传失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "文件上传成功",
	})
}
