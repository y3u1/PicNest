package initializer

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Engine *gorm.DB

func InitializeDB() error {

	Engine, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	return err
}
