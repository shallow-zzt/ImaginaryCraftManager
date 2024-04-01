package weblogin

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/go-ini/ini"
)

// 生成随机字符串
func generateRandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	//rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenLoginToken(filePath string) {
	username := generateRandomString(16)
	password := generateRandomString(16)
	rconPassword := generateRandomString(16)

	// 创建INI文件
	cfg := ini.Empty()
	section, err := cfg.NewSection("User")
	if err != nil {
		fmt.Printf("创建失败: %v\n", err)
		os.Exit(1)
	}

	section.ReflectFrom(&struct{ Username string }{username})
	section.ReflectFrom(&struct{ Password string }{password})
	section.ReflectFrom(&struct{ Rcon_Password string }{rconPassword})

	err = cfg.SaveTo(filePath)
	if err != nil {
		fmt.Printf("INI保存失败: %v\n", err)
		os.Exit(1)
	}
	//fmt.Println(username, password)
}
