package services

import (
	"PicNest/internal/model"
	"PicNest/internal/repository"
	"PicNest/internal/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	AuthRepo *repository.AuthRepository
}

func NewAuthService(engine *gorm.DB) *AuthService {
	return &AuthService{AuthRepo: repository.NewAuthRepository(engine)}
}
func (as *AuthService) QueryRecord(username string) (model.UserLoginInfo, error) {
	return as.AuthRepo.QueryUser(username)
}
func (as *AuthService) QueryToken(token string) (model.UserLoginInfo, error) {
	return as.AuthRepo.QueryToken(token)
}
func (as *AuthService) CreateRecord(username string) (model.UserLoginInfo, error) {
	token, err := utils.GenerateToken(username)
	if err != nil {
		return model.UserLoginInfo{}, err
	}
	return as.AuthRepo.CreateUserLoginInfo(username, token)
}
func (as *AuthService) UpdateRecord(username string) (model.UserLoginInfo, error) {
	token, err := utils.GenerateToken(username)
	if err != nil {
		return model.UserLoginInfo{}, err
	}
	return as.AuthRepo.UpdateLoginInfo(username, token)
}
func (as *AuthService) Authenticate(token string) error {
	ul, err := as.AuthRepo.QueryToken(token)
	if err != nil {
		return err
	}
	err = utils.ValidateToken(ul.Token)
	if err != nil {
		return err
	}
	return nil
}
