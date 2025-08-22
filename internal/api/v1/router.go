package v1

import (
	"PicNest/internal/controller"
	"PicNest/internal/middleware"
	"PicNest/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRouter(r *gin.Engine, engine *gorm.DB) {

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("static/*")
	// r.Static("./internal/static")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{})
	})
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(200, "upload.html", gin.H{})
	})
	AuthService := services.NewAuthService(engine)
	UserService := services.NewUserService(engine)
	UserController := controller.NewUserController(UserService, AuthService)
	UploadController := controller.NewUploadController()
	v1 := r.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/login", UserController.Login)
			user.POST("/register", UserController.Register)
		}
		v1.POST("/upload", middleware.Auth(AuthService), UploadController.Upload)

	}
}
