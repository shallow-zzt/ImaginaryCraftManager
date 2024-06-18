package logger

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	NewLogger("debug")
	Debugln("测试")
	Debugf("测试%v", 1)
}
