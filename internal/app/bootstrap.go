package app

import (
	v1 "PicNest/internal/api/v1"
	"PicNest/internal/app/config"
	"PicNest/internal/app/initializer"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Start 启动服务
func Start() {
	err := config.LoadConfig()
	if err != nil {
		log.Error("配置文件加载错误: %v", err)
		return
	}

	err = InitializeAll()
	if err != nil {
		log.Error("模块初始化错误: %v", err)
		return
	}

	r := gin.New()
	v1.SetRouter(r, initializer.Engine)

	err = r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		log.Error("服务启动错误: %v", err)
		return
	}
}

// InitializeAll 初始化所有模块
func InitializeAll() error {
	err := initializer.InitializeLogger()
	if err != nil {
		return fmt.Errorf("日志初始化错误: %v", err)
	}
	err = initializer.InitializeDB()
	if err != nil {
		return fmt.Errorf("Sqlite初始化错误: %v", err)
	}
	return nil
}
