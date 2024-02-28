package serverCmd

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestServerCmd(t *testing.T) {
	serverPath := "..\\..\\fabric-server"
	cmd := exec.Command("cmd.exe", "/c", "start.bat")
	cmd.Dir = serverPath

	// 启动命令并记录输出
	javaPid, err := CmdRecording(cmd)
	fmt.Println(javaPid)
	if err != nil {
		return
	}

	// 在这里执行一些其他操作，然后当需要时关闭命令及其管道
	fmt.Println("Press enter to close the process and pipe...")
	fmt.Scanln()

	err = CloseProcessAndPipe(cmd)
	if err != nil {
		fmt.Println("Error closing process and pipe:", err)
	}
}
