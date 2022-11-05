package tests

import (
	"gin-message-board/models"
	"gin-message-board/tools"
	"testing"
)

// TestUsernameAvailability 测试用户名合法性
func TestUsernameAvailability(t *testing.T) {
	tools.SaveLists()

	if !models.IsUsernameAvailable("newuser") {
		t.Fail()
	}

	if models.IsUsernameAvailable("user1") {
		t.Fail()
	}

	models.RegisterNewUser("newuser", "newpass")

	if models.IsUsernameAvailable("newuser") {
		t.Fail()
	}

	tools.RestoreLists()
}

// TestValidUserRegistration 测试合格的用户注册
func TestValidUserRegistration(t *testing.T) {
	tools.SaveLists()

	u, err := models.RegisterNewUser("newuser", "newpass")

	if err != nil || u.Username == "" {
		t.Fail()
	}

	tools.RestoreLists()
}

// TestInvalidUserRegistration 测试不合格的用户注册
func TestInvalidUserRegistration(t *testing.T) {
	tools.SaveLists()

	u, err := models.RegisterNewUser("user1", "pass1")

	if err == nil || u != nil {
		t.Fail()
	}

	u, err = models.RegisterNewUser("newuser", "")

	if err == nil || u != nil {
		t.Fail()
	}

	tools.RestoreLists()
}

// TestUserValidity 测试用户合法性
func TestUserValidity(t *testing.T) {
	if !models.IsUserValid("user1", "pass1") {
		t.Fail()
	}

	if models.IsUserValid("user2", "pass1") {
		t.Fail()
	}

	if models.IsUserValid("user1", "") {
		t.Fail()
	}

	if models.IsUserValid("", "pass1") {
		t.Fail()
	}

	if models.IsUserValid("User1", "pass1") {
		t.Fail()
	}
}
