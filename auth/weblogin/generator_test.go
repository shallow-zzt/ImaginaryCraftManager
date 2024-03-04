package weblogin

import (
	"fmt"
	"testing"
)

func TestGenLoginToken(t *testing.T) {
	iniPath := "authorities.ini"
	LoadUsers(iniPath)
	fmt.Println(GetUsername(iniPath))
}
