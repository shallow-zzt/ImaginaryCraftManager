package weblogin

import (
	"net/http"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func CheckLogin(storedUser string, storedPassword string) bool {
	return storedPassword == users[storedUser] && storedPassword != "" && storedUser != ""
}

// 检查用户是否已经登录的中间件
func CheckIsLogined(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("session")
	if err != nil || cookie.Value == "" || users[cookie.Value] == "" {
		return false
	}
	return true
}
