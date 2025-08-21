package services

import (
	"PicNest/internal/model"
	"PicNest/internal/repository"

	"gorm.io/gorm"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(engine *gorm.DB) *UserService {
	return &UserService{UserRepo: repository.NewUserRepository(engine)}
}
func (us *UserService) Query(username string) (model.UserInfo, error) {
	return us.UserRepo.QueryUser(username)
}

func (us *UserService) Register(username string, password string) (model.UserInfo, error) {
	return us.UserRepo.CreateUser(username, password)
}

func (us *UserService) Deregister(username string) error {
	return us.UserRepo.DeleteUser(username)
}
