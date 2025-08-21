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
