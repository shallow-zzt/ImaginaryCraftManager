package jwtApiToken

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secret = "ImaginaryToken"

func GenerateJwtToken(username string) string {
	// 生成 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // 过期时间为 1 小时
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
		return "-1"
	}
	return tokenString
}

func TokenTester(w http.ResponseWriter, r *http.Request) bool {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		fmt.Fprintf(w, "缺少 token")
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Fprintf(w, "解析 token 失败: %v", err)
		return false
	}

	if token.Valid {
		return true
	} else {
		fmt.Fprintf(w, "token 无效")
		return false
	}
}
