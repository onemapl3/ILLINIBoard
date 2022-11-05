package routers

import (
	"gin-message-board/controllers"
	auth "gin-message-board/middlewares"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes 路由配置
func InitializeRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(auth.SetUserStatus())
	router.LoadHTMLGlob("templates/*")
	// 主页路由
	router.GET("/", controllers.ShowIndexPage)
	// 留言路由组
	articleRoutes := router.Group("/message")
	{
		// route from Part 1 of the tutorial
		articleRoutes.GET("/view/:message_id", controllers.GetMessage)

		articleRoutes.GET("/create", controllers.ShowMessageCreationPage)

		articleRoutes.POST("/create", controllers.CreateMessage)
	}

	// 用户路由组
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", controllers.ShowRegistrationPage)
		userRoutes.POST("/register", controllers.Register)

		userRoutes.GET("/login", controllers.ShowLoginPage)
		userRoutes.POST("/login", controllers.PerformLogin)
		userRoutes.GET("/logout", controllers.Logout)
	}
	return router
}
