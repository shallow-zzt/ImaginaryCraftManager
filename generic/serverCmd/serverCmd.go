package serverCmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type CommandManager struct {
	cmd    *exec.Cmd
	stdout *bufio.Scanner
}

func SetCmdParameter(serverDir string, serverMemory string) error {
	serverRunCommand := "java -Xmx" + serverMemory + "G -jar fabric-server-launch.jar nogui"
	cmdFileName := serverDir + "\\start.bat"

	file, err := os.Create(cmdFileName)
	if err != nil {
		fmt.Println("cmd运行脚本创建失败:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(serverRunCommand)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}

func NewCmdManager(serverDir string) (*CommandManager, error) {
	cmd := exec.Command("cmd.exe", "/c", "start.bat")
	cmd.Dir = serverDir
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("获取进程管道失败:", err)
		return nil, err
	}
	return &CommandManager{
		cmd:    cmd,
		stdout: bufio.NewScanner(stdout),
	}, nil
}

func CloseProcessAndPipe(cm *CommandManager) error {
	//暴力关闭Java
	//因为我也没想到更好的办法
	//我不清楚这样做，会不会把所有需要java运行的程序都关了 ^_^

	//经过测试，至少客户端和服务端同时启动时，不会关闭客户端 ^_^
	exec.Command("taskkill", "/f", "/im", "java.exe").Run()
	if err := cm.cmd.Process.Kill(); err != nil {
		fmt.Println("进程关闭失败:", err)
		return err
	}

	return nil
}

func CmdRecording(cm *CommandManager) (javaPID int, err error) {
	var outputLines []string

	if err := cm.cmd.Start(); err != nil {
		fmt.Println("服务器启动失败:", err)
		return 0, err
	}

	go func() {
		for cm.stdout.Scan() {
			outputLines = append(outputLines, cm.stdout.Text())
			fmt.Println(cm.stdout.Text())
		}
	}()

	javaPID = cm.cmd.Process.Pid
	fmt.Println(javaPID)

	return javaPID, nil
}
