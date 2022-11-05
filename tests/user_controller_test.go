package tests

import (
	"gin-message-board/controllers"
	"gin-message-board/tools"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

// getLoginPOSTPayload
// 获取登陆时的表单信息
func getLoginPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "user1")
	params.Add("password", "pass1")

	return params.Encode()
}

// getRegistrationPOSTPayload
// 获取注册时表单信息
func getRegistrationPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "u1")
	params.Add("password", "p1")

	return params.Encode()
}

// TestShowRegistrationPageUnauthenticated
// 测试未认证的展示注释页面
func TestShowRegistrationPageUnauthenticated(t *testing.T) {
	r := tools.GetRouter(true)

	r.GET("/u/register", controllers.ShowRegistrationPage)

	req, _ := http.NewRequest("GET", "/u/register", nil)

	tools.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Register</title>") > 0

		return statusOK && pageOK
	})
}

// TestRegisterUnauthenticated
// 测试未认证的注册页面
func TestRegisterUnauthenticated(t *testing.T) {
	tools.SaveLists()
	w := httptest.NewRecorder()

	r := tools.GetRouter(true)

	r.POST("/u/register", controllers.Register)

	registrationPayload := getRegistrationPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/register", strings.NewReader(registrationPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Successful registration &amp; Login</title>") < 0 {
		t.Fail()
	}
	tools.RestoreLists()
}

// TestRegisterUnauthenticatedUnavailableUsername
// 测试注册未认证的用户名
func TestRegisterUnauthenticatedUnavailableUsername(t *testing.T) {
	tools.SaveLists()
	w := httptest.NewRecorder()

	r := tools.GetRouter(true)

	r.POST("/u/register", controllers.Register)

	registrationPayload := getLoginPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/register", strings.NewReader(registrationPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
	tools.RestoreLists()
}

// TestShowLoginPageUnauthenticated
// 测试展示未认证的登录页面
func TestShowLoginPageUnauthenticated(t *testing.T) {
	r := tools.GetRouter(true)

	r.GET("/u/login", controllers.ShowLoginPage)

	req, _ := http.NewRequest("GET", "/u/login", nil)

	tools.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Login</title>") > 0

		return statusOK && pageOK
	})
}

// TestLoginUnauthenticated
// 测试未认证登录
func TestLoginUnauthenticated(t *testing.T) {
	tools.SaveLists()
	w := httptest.NewRecorder()
	r := tools.GetRouter(true)

	r.POST("/u/login", controllers.PerformLogin)

	loginPayload := getLoginPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Successful Login</title>") < 0 {
		t.Fail()
	}
	tools.RestoreLists()
}

// TestLoginUnauthenticatedIncorrectCredentials
// 测试未认证不正确的登录证书
func TestLoginUnauthenticatedIncorrectCredentials(t *testing.T) {
	tools.SaveLists()
	w := httptest.NewRecorder()
	r := tools.GetRouter(true)

	r.POST("/u/login", controllers.PerformLogin)

	loginPayload := getRegistrationPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
	tools.RestoreLists()
}
