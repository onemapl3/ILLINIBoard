package controllers

import (
	"gin-message-board/database"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GenerateSessionToken 生成随机的16字符字符串作为会话标记(生产环境不应使用这种方式)
func GenerateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

// ShowRegistrationPage 展示注册页面
func ShowRegistrationPage(c *gin.Context) {
	render(c, gin.H{"title": "注册"}, "register.html")
}

// Register 注册页面逻辑
func Register(c *gin.Context) {
	// 从Post表单获取name="username"和name="password"
	username := c.PostForm("username")
	password := c.PostForm("password")

	err := database.RegisterNewUser(username, password)
	if err != nil {
		// 如果用户名或密码不合法展示错误再登录界面
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "注册失败",
			"ErrorMessage": err.Error()})
	}
	token := GenerateSessionToken()
	// 如果创建了用户，需要在cookie中设置token，然后登录
	c.SetCookie("token", token, 3600, "", "", false, true)
	// 设置"is_logged_"=true
	c.Set("is_logged_in", true)

	render(c, gin.H{"title": "成功注册，登录成功"}, "login-successful.html")

}

// ShowLoginPage 展示登录界面
func ShowLoginPage(c *gin.Context) {
	render(c, gin.H{
		"title": "登录",
	}, "login.html")
}

// PerformLogin 执行登录逻辑
func PerformLogin(c *gin.Context) {
	// 从Post表单获取name="username"和name="password"
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 判断是否合法
	valid, err := database.IsUserValid(username, password)
	if err != nil || valid {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "登录失败",
			"ErrorMessage": "Invalid credentials provided"})
	}
	// 生成token并设置
	token := GenerateSessionToken()
	c.SetCookie("token", token, 3600, "", "", false, true)
	render(c, gin.H{
		"title": "成功登录"}, "login-successful.html")

}

// Logout 用户登出逻辑
func Logout(c *gin.Context) {
	// 设置Cookie token令牌
	c.SetCookie("token", "", -1, "", "", false, true)
	// 重定向到首页
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
