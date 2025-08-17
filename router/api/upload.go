package api

import (
	"PicNext/utils"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

var whitelist = [...]string{"png", "jpg"}

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "文件上传失败！",
		})
	}
	file.Header.Get()
	filename := utils.RandName()
	filetype, err := detectFileType(file.)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "不允许的文件类型",
		})
	}
	filetype = strings.Split(filetype, "/")[1]
	filename = filename + "." + filetype
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

func detectFileType(f *http.File) (string,error){
	mime , err := mimetype.DetectReader(f)
	if err != nil {
		return "",err
	}
	return mime.String(),nil
}
