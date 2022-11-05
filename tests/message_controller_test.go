package tests

import (
	"gin-message-board/controllers"
	"gin-message-board/tools"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestShowIndexPageUnauthenticated 测试对主页的匿名GET请求是否返回主页
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := tools.GetRouter(true)

	r.GET("/", controllers.ShowIndexPage)

	// 创建一个请求以发送到上述路由
	req, _ := http.NewRequest("GET", "/", nil)

	tools.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// 测试http状态码是否为200
		statusOK := w.Code == http.StatusOK

		// 测试页面标题是否为“主页”
		// 在此可以使用解析和处理HTML页面的库来执行更多详细的测试
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>主页</title>") > 0

		return statusOK && pageOK
	})
}
