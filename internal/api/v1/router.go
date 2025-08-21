package v1

import (
	v1 "PicNest/internal/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRouter(r gin.Engine, Engine *gorm.DB) {

	// r.Static("./internal/static")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	})
	r.POST("/upload", v1.UploadImage)
}
