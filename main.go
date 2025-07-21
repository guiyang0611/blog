package main

import (
	"blog/config"
	"blog/log"
	"blog/models"
	"blog/routers"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 1. 初始化日志
	log.InitLogger()
	// 2. 释放日志
	defer log.Logger.Sync()
	// 3. 初始化数据库
	err := config.InitDb()
	if err != nil {
		log.Logger.Fatal("数据库初始化失败", zap.Error(err))
	}
	db := config.GormDB
	// 5. 自动迁移
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	// 6. 初始化路由
	router := routers.SetupRouter()
	router.Use(utils.ErrorHandler())
	router.Use(gin.Logger())
	// 7. 启动服务
	port := config.Cfg.Server.Port
	log.Logger.Info("服务启动", zap.String("port", port))
	err = router.Run(port)
	if err != nil {
		log.Logger.Fatal("服务启动失败", zap.Error(err))
	}
}
