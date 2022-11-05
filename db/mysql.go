package db

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"illini-board/config"
	"illini-board/models"
	"strings"
)

var db *gorm.DB

// 初始化連線資料庫
func InitMySql() (err error) {
	var c *config.Conf

	//獲取yaml配置引數
	conf := c.GetConf()

	//將yaml配置引數拼接成連線資料庫的url
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DbName,
	)

	//連線資料庫
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

// GetDB 返回数据库指针
func GetDB() *gorm.DB {
	return db
}

// GetAllMessages 从数据库拿取所有留言
func GetAllMessages() ([]models.Message, error) {
	var messages []models.Message
	result := db.Find(&messages)
	return messages, result.Error
}

// GetMessageByID 通过ID拿取信息
func GetMessageByID(id int) (models.Message, error) {
	var message models.Message
	result := db.Where("ID = ?", id).First(&message)
	return message, result.Error
}

// CreateNewMessage 将新的留言写入数据库
func CreateNewMessage(title, content string) (models.Message, error) {
	message := models.Message{Title: title, Content: content}
	result := db.Create(&message)
	return message, result.Error
}

// RegisterNewUser 注册新用户写入数据库
func RegisterNewUser(username, password string) error {
	// 判断密码是否为空
	if strings.TrimSpace(password) == "" {
		return errors.New("Password should not be empty")
	}
	// 判断用户名是否可用
	userAvailable, err := IsUsernameAvailable(username)
	if err == nil && !userAvailable {
		return errors.New("Username Unavailable")
	}
	// 写入数据库
	user := models.User{Username: username, Password: password}
	result := db.Create(&user)
	return result.Error
}

// IsUsernameAvailable 判断用户名是否可用
func IsUsernameAvailable(username string) (bool, error) {
	var user []models.User
	result := db.Where("Username = ?", username).Find(&user)
	return result.RowsAffected == 0, result.Error
}

// IsUserValid 判断用户是否存在,合法
func IsUserValid(username, password string) (bool, error) {
	var user []models.User
	result := db.Where("Username = ? AND Password = ?", username, password).Find(&user)
	return result.RowsAffected == 1, result.Error
}
