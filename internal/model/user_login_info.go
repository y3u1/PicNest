package model

import "gorm.io/gorm"

type UserLoginInfo struct {
	gorm.Model
	Username string `gorm:"notnull"`
	Token    string `gorm:"notnull"`
}
