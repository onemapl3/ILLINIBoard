package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

// User 用户结构体
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"-"`
}

/*
	硬编码方式返回留言
*/
var UserList = []User{
	User{Username: "用户1", Password: "yonghu11111"},
	User{Username: "用户2", Password: "yonghu22222"},
}

// IsUserValid 检查用户名和密码
func IsUserValid(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

// 注册新用户
func RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("密码不能为空")
	} else if !IsUsernameAvailable(username) {
		return nil, errors.New("用户名不可用")
	}

	u := User{Username: username, Password: password}

	UserList = append(UserList, u)

	return &u, nil
}

// 检查提供的用户名是否可用
func IsUsernameAvailable(username string) bool {
	for _, u := range UserList {
		if u.Username == username {
			return false
		}
	}
	return true
}
