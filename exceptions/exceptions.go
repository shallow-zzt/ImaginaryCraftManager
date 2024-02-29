package exceptions

import "fmt"

// 错误类型错误码
type ErrorWithCode struct {
	Code    int
	Message string
}

func (e ErrorWithCode) ErrorWithCode() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
