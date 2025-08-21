package controller

import (
	"PicNest/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService *services.UserService
	authService *services.AuthService
}
type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserInfoDto struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func NewUserController(UserService *services.UserService, authService *services.AuthService) *UserController {
	return &UserController{userService: UserService, authService: authService}
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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password)); err == nil {
		//账号存在,一定有记录,直接更新
		userInfo, err := uc.authService.UpdateRecord(userDto.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"success": false,
			})
			return
		}
		c.Header("Authorization", userInfo.Token)
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
			"success": true,
			"id":      user.ID,
			"token":   userInfo.Token,
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
	//注册成功,创建登录记录
	userInfo, err := uc.authService.CreateRecord(userDto.Username)
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
	c.Header("Authorization", userInfo.Token)

}
