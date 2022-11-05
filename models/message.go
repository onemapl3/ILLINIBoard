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
}

/*
	硬编码方式返回留言
*/
var MessageList = []Message{
	Message{ID: 1, Title: "留言标题1", Content: "留言内容1"},
	Message{ID: 2, Title: "留言标题2", Content: "留言内容2"},
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
