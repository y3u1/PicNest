package repository

import (
	"PicNest/internal/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	engine *gorm.DB
}

func NewUserRepository(engine *gorm.DB) *UserRepository {
	return &UserRepository{engine: engine}
}

func (r *UserRepository) QueryUser(username string) (model.UserInfo, error) {
	var user model.UserInfo
	result := r.engine.First(&user, "username = ?", username)
	if result.RowsAffected == 0 {
		return user, result.Error
	}
	return user, nil
}
func (r *UserRepository) CreateUser(username string, password string) (model.UserInfo, error) {
	var user model.UserInfo
	result := r.engine.First(&user, "username = ?", username)
	if result.RowsAffected == 0 {
		return user, errors.New("用户已存在")
	}
	user = model.UserInfo{Username: username, Password: password}
	result = r.engine.Create(&user)
	return user, result.Error
}
func (r *UserRepository) DeleteUser(username string) error {
	var user model.UserInfo
	result := r.engine.First(&user, "username = ?", username)
	if result.RowsAffected == 0 {
		return errors.New("不存在该用户")
	}
	r.engine.Delete(user)
	return result.Error
}
