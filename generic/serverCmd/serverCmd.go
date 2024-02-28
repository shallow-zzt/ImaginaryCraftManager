package serverCmd

import (
	"bufio"
	"fmt"
	"os/exec"
)

func cmdRecording(serverPath string) {
	cmd := exec.Command("cmd.exe", "/c", "start.bat")
	cmd.Dir = serverPath

	// 创建一个管道来获取命令的标准输出和标准错误输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe:", err)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Error creating StderrPipe:", err)
		return
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	// 从标准输出和标准错误输出读取命令的输出
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// 等待命令执行完成
	if err := cmd.Wait(); err != nil {
		fmt.Println("Command finished with error:", err)
	}
}
