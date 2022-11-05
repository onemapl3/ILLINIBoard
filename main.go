package main

import (
	"gin-message-board/config"
	"gin-message-board/database"
	"gin-message-board/routers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	config.Init()
	database.Init()
	var router = routers.InitializeRoutes()

	// 启动服务
	router.Run()

}
