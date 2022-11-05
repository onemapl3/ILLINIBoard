package tests

import (
	"gin-message-board/models"

	"testing"
)

// TestGetAllMessages 测试获取所有留言
func TestGetAllMessages(t *testing.T) {
	mlist := models.GetAllMessages()

	// 检查返回的条目列表的长度是否为与保存链表的全局变量的长度相同
	if len(mlist) != len(models.MessageList) {
		t.Fail()
	}

	// 检查每个message是否相同
	for i, v := range mlist {
		if v.Content != models.MessageList[i].Content ||
			v.ID != models.MessageList[i].ID ||
			v.Title != models.MessageList[i].Title {

			t.Fail()
			break
		}
	}
}

// TestGetMessageByID 根据留言的ID测试获取留言的函数
func TestGetMessageByID(t *testing.T) {
	m, err := models.GetMessageByID(1)

	if err != nil || m.ID != 1 || m.Title != "留言标题1" || m.Content != "留言内容1" {
		t.Fail()
	}
}
