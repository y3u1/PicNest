package model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Username string `gorm:"notnull";json:"username"`
	Password string `gorm:"notnull";json:"password"`
}

// func (u UserInfo) TableName() string {
// 	return "userinfo"
// }
