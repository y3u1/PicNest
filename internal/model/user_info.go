package model

import (
	"PicNest/internal/utils"
	"strings"

	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	Username string `gorm:"notnull";json:"username"`
	Password string `gorm:"notnull";json:"password"`
}

func (u *UserInfo) BeforeSave(*gorm.DB) error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	u.Username = strings.TrimSpace(u.Username)
	return nil
}
