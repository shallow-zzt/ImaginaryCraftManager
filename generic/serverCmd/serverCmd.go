package serverCmd

import (
	"bufio"
	"fmt"
	"os/exec"
)

func CloseProcessAndPipe(cmd *exec.Cmd) error {
	if err := cmd.Process.Kill(); err != nil {
		fmt.Println("进程关闭失败:", err)
		return err
	}
	return nil
}

func CmdRecording(cmd *exec.Cmd) (javaPID int, err error) {
	var outputLines []string
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("获取进程管道失败:", err)
		return 0, err
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("指令执行失败:", err)
		return 0, err
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			outputLines = append(outputLines, scanner.Text())
			fmt.Println(scanner.Text())
		}
	}()

	javaPID = cmd.Process.Pid
	fmt.Println(javaPID)

	if err := cmd.Wait(); err != nil {
		fmt.Println("指令执行失败:", err)
		return javaPID, err
	}
	return javaPID, nil
}
