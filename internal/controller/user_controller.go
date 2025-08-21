package controller

import (
	"PicNest/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}
type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserController(UserService *services.UserService) *UserController {
	return &UserController{userService: UserService}
}

func (uc *UserController) Login(c *gin.Context) {
	var userDto UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求数据格式错误",
		})
		return
	}
	user, err := uc.userService.Query(userDto.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if userDto.Password == user.Password {
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
			"success": true,
			"id":      user.ID,
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "不正确的用户名或者密码",
			"success": false,
		})
	}

}
func (uc *UserController) Register(c *gin.Context) {
	var userDto UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求数据格式错误",
			"success": false,
		})
		return
	}
	user, err := uc.userService.Register(userDto.Username, userDto.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"success": true,
		"id":      user.ID,
	})

}
