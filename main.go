package main

import (
	"illini-board/db"
	"illini-board/routers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	if err := db.InitMySql(); err != nil {
		panic(err)
	}

	var router = routers.InitializeRoutes()

	// 启动服务
	router.Run()

}
