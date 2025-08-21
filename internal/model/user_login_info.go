package model

import "gorm.io/gorm"

type UserLoginInfo struct {
	gorm.Model
	Token string
}

func (u UserLoginInfo) TableName() string {
	return "userlogininfo"
}
