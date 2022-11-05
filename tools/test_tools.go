// test_tools.go
// 辅助测试函数

package tools

import (
	auth "gin-message-board/middlewares"
	"gin-message-board/models"

	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpUserList []models.User
var tmpMessageList []models.Message

//  TestMain 在执行测试函数之前进行setup
func TestMain(m *testing.M) {
	//设置gin为测试模式
	gin.SetMode(gin.TestMode)

	// 运行其他测试
	os.Exit(m.Run())
}

// GetRouter 在测试期间创建getRouter函数
func GetRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
		r.Use(auth.SetUserStatus())
	}
	return r
}

// TestHTTPResponse 处理请求并测试其响应的函数
func TestHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// 创建响应记录器
	w := httptest.NewRecorder()

	// 创建service并处理上述请求
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// SaveLists 这个函数用于将主列表存储到临时列表中进行测试
func SaveLists() {
	tmpUserList = models.UserList
	tmpMessageList = models.MessageList
}

// RestoreLists 此函数用于从临时列表恢复主列表
func RestoreLists() {
	models.UserList = tmpUserList
	models.MessageList = tmpMessageList
}

// TestMiddlewareRequest 测试中间件请求
func TestMiddlewareRequest(t *testing.T, r *gin.Engine, expectedHTTPCode int) {
	req, _ := http.NewRequest("GET", "/", nil)

	TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == expectedHTTPCode
	})
}
