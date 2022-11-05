package models

import (
	"errors"

	"gorm.io/gorm"
)

// Message 留言结构
type Message struct {
	gorm.Model
	ID      int    `gorm : NOT NULL autoIncrement json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Tag     string `json:"tag"`
	Author  string `json:"author"`
	Board   string `json:"board"`
}

/*
硬编码方式返回留言
*/
var MessageList = []Message{
	Message{ID: 1, Title: "留言标题1", Content: "留言内容1", Date: "2022-11-05:06:36:01", Tag: "Tag1", Author: "user1", Board: "1"},
	Message{ID: 2, Title: "留言标题2", Content: "留言内容2", Date: "2022-11-05:06:36:01", Tag: "Tag2", Author: "user2", Board: "1"},
}

// GetAllMessages 返回留言列表
func GetAllMessages() []Message {
	return MessageList
}

// GetMessageByID 根据提供的ID获取一个留言
func GetMessageByID(id int) (*Message, error) {
	for _, m := range MessageList {
		if m.ID == id {
			return &m, nil
		}
	}
	return nil, errors.New("Message not found")
}
