package initializer

import (
	"PicNest/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Engine *gorm.DB

func InitializeDB() error {

	engine, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	exist := engine.Migrator().HasTable("user_infos")
	if !exist {
		engine.AutoMigrate(&model.UserInfo{})
	}
	exist = engine.Migrator().HasTable("user_login_infos")
	if !exist {
		engine.AutoMigrate(&model.UserLoginInfo{})
	}
	Engine = engine
	return err
}
