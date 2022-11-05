package controllers

import (
	"fmt"
	"gin-message-board/database"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ShowIndexPage 从数据库读取所有留言并传到templates/index.html
func ShowIndexPage(c *gin.Context) {
	messages, err := database.GetAllMessages()
	if err != nil {
		panic(fmt.Errorf("数据库读取留言错误: %s \n", err))
	}
	render(c, gin.H{
		"title":   "主页",
		"payload": messages}, "index.html",
	)

}

// GetMessage 从上下文获取信息
func GetMessage(c *gin.Context) {
	// 登录接口
	loggedInInterface, _ := c.Get("is_logged_in")

	// 将字符串messageID转换成int
	messageID, err := strconv.Atoi(c.Param("message"))
	if err != nil {
		// 如果在URL中指定了无效的留言ID，则终止并显示错误
		c.AbortWithStatus(http.StatusNotFound)
	}

	// 检查留言在数据库中是否存在
	message, err := database.GetMessageByID(messageID)
	if err != nil {
		panic(fmt.Errorf("留言在数据库中不存在：%s \n", c.AbortWithError(http.StatusNotFound, err).Error()))
	}
	c.HTML(
		http.StatusOK,
		// 使用"message.html"模板
		"message.html",
		// 传递页面使用的数据
		gin.H{
			"title":        message.Title,
			"payload":      message,
			"is_logged_in": loggedInInterface.(bool),
		},
	)
}

// render
func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// 响应JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// 响应XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// 默认响应HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}

// ShowMessageCreationPage 展示留言创建页面
func ShowMessageCreationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "创建新留言"}, "create-message.html")
}

// CreateMessage 留言提交页面
func CreateMessage(c *gin.Context) {
	// Post表单获取name="title"和name="content"
	title := c.PostForm("title")
	content := c.PostForm("content")
	// 将新的留言写入数据库
	m, err := database.CreateNewMessage(title, content)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	render(c, gin.H{
		"title":   "提交成功",
		"payload": m}, "submission-successful.html")

}
