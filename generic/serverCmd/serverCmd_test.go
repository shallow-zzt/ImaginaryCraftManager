package serverCmd

import (
	"fmt"
	"testing"
	"time"
)

func TestServerSetMemory(t *testing.T) {
	err := SetCmdParameter("..\\..\\fabric-server", 3)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func TestServerCmd(t *testing.T) {
	manager, err := NewCmdManager("..\\..\\fabric-server")
	if err != nil {
		fmt.Println("cmd管道创建失败:", err)
		return
	}

	//两次启动进程测试
	javaPid, err := CmdRecording(manager)
	fmt.Println(javaPid)
	if err != nil {
		return
	}

	time.Sleep(5 * time.Second)
	fmt.Println("程序第一次启动完毕，开始终止进程……")

	err = CloseProcessAndPipe(manager)
	if err != nil {
		fmt.Println("Error closing process and pipe:", err)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("准备第二次开启进程……")

	manager, err = NewCmdManager("..\\..\\fabric-server")
	if err != nil {
		fmt.Println("cmd管道创建失败:", err)
		return
	}

	javaPid, err = CmdRecording(manager)
	fmt.Println(javaPid)
	if err != nil {
		return
	}
	time.Sleep(5 * time.Second)
	fmt.Println("第二次启动进程关闭……")

	err = CloseProcessAndPipe(manager)
	if err != nil {
		fmt.Println("Error closing process and pipe:", err)
	}
}
