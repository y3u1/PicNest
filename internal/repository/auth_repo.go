package repository

import (
	"PicNest/internal/model"

	"gorm.io/gorm"
)

type AuthRepository struct {
	engine *gorm.DB
}

func NewAuthRepository(engine *gorm.DB) *AuthRepository {
	return &AuthRepository{engine: engine}
}

func (r *AuthRepository) QueryUser(username string) (model.UserLoginInfo, error) {
	var userlogin model.UserLoginInfo

	result := r.engine.First(&userlogin, "username = ?", username)
	if result.RowsAffected == 0 {
		return userlogin, result.Error
	}
	return userlogin, nil
}
func (r *AuthRepository) CreateUserLoginInfo(username, token string) (model.UserLoginInfo, error) {
	var userlogin model.UserLoginInfo

	result := r.engine.First(&userlogin, "username = ?", username)
	if result.RowsAffected != 0 {
		return r.UpdateLoginInfo(username, token)
	}
	userlogin = model.UserLoginInfo{Username: username, Token: token}

	result = r.engine.Create(&userlogin)
	return userlogin, result.Error
}
func (r *AuthRepository) UpdateLoginInfo(username, token string) (model.UserLoginInfo, error) {
	var userlogin model.UserLoginInfo
	userlogin.Token = token
	result := r.engine.Save(&userlogin)
	return userlogin, result.Error
}
