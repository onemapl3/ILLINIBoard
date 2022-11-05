package tests

import (
	auth "gin-message-board/middlewares"
	"gin-message-board/tools"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestEnsureLoggedInUnauthenticated 测试确认登出
func TestEnsureLoggedInUnauthenticated(t *testing.T) {
	r := tools.GetRouter(false)
	r.GET("/", setLoggedIn(false), auth.EnsureLoggedIn(), func(c *gin.Context) {
		t.Fail()
	})

	tools.TestMiddlewareRequest(t, r, http.StatusUnauthorized)
}

// setLoggedIn 设置"is_logged_in"
func setLoggedIn(b bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("is_logged_in", b)
	}
}

// TestEnsureLoggedInAuthenticated 测试确认登录认证
func TestEnsureLoggedInAuthenticated(t *testing.T) {
	r := tools.GetRouter(false)
	r.GET("/", setLoggedIn(true), auth.EnsureLoggedIn(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	tools.TestMiddlewareRequest(t, r, http.StatusOK)
}

// TestEnsureNotLoggedInAuthenticated 测试未登录认证
func TestEnsureNotLoggedInAuthenticated(t *testing.T) {
	r := tools.GetRouter(false)
	r.GET("/", setLoggedIn(true), auth.EnsureNotLoggedIn(), func(c *gin.Context) {
		t.Fail()
	})

	tools.TestMiddlewareRequest(t, r, http.StatusUnauthorized)
}

// TestEnsureNotLoggedInAuthenticated 测试未登录认证
func TestEnsureNotLoggedInUnauthenticated(t *testing.T) {
	r := tools.GetRouter(false)
	r.GET("/", setLoggedIn(false), auth.EnsureNotLoggedIn(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	tools.TestMiddlewareRequest(t, r, http.StatusOK)
}

// TestSetUserStatusAuthenticated 测试用户状态
func TestSetUserStatusAuthenticated(t *testing.T) {
	r := tools.GetRouter(false)
	r.GET("/", auth.SetUserStatus(), func(c *gin.Context) {
		loggedInInterface, exists := c.Get("is_logged_in")
		if !exists || !loggedInInterface.(bool) {
			t.Fail()
		}
	})

	w := httptest.NewRecorder()

	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}

	r.ServeHTTP(w, req)
}

// TestSetUserStatusUnauthenticated 测试设置用户状态
func TestSetUserStatusUnauthenticated(t *testing.T) {
	r := tools.GetRouter(false)
	r.GET("/", auth.SetUserStatus(), func(c *gin.Context) {
		loggedInInterface, exists := c.Get("is_logged_in")
		if exists && loggedInInterface.(bool) {
			t.Fail()
		}
	})

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)
}
