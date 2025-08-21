package model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u UserInfo) TableName() string {
	return "userinfo"
}
