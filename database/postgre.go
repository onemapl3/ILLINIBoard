package database

import (
	"errors"
	"fmt"
	"gin-message-board/models"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init 初始化postgresql数据库设置
func Init() {
	var dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		viper.GetString("database.server"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"),
		viper.GetInt("database.ports"),
	)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Fatal database error: \n", err))
	}

	err = db.AutoMigrate(&models.Message{}, &models.User{})
	if err != nil {
		panic(err)
	}
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
		return errors.New("密码不能为空")
	}
	// 判断用户名是否可用
	userAvailable, err := IsUsernameAvailable(username)
	if err == nil && !userAvailable {
		return errors.New("用户名不可用")
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
