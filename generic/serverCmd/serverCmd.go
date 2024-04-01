package serverCmd

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/gorilla/websocket"
)

type CommandManager struct {
	cmd    *exec.Cmd
	stdout *bufio.Scanner
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func SetCmdParameter(serverDir string, serverMemory int) error {

	serverRunCommand := "java -Xmx" + strconv.Itoa(serverMemory) + "G -jar fabric-server-launch.jar nogui"
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

func GetCmdParameter(serverDir string) string {
	memoryRegex := "-Xmx(.*)G"
	cmdFileName := serverDir + "\\start.bat"
	file, err := os.Open(cmdFileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	re := regexp.MustCompile(memoryRegex)
	matches := re.FindAllStringSubmatch(string(buf[:n]), 1)[0][1]
	return matches
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

	if err := cm.cmd.Process.Kill(); err != nil {
		fmt.Println("进程关闭失败:", err)
		exec.Command("taskkill", "/f", "/im", "java.exe").Run()
		return err
	}
	fmt.Println("关闭成功")
	exec.Command("taskkill", "/f", "/im", "java.exe").Run()

	return nil
}

func CmdRecording(w http.ResponseWriter, r *http.Request, cm *CommandManager) (javaPID int, err error) {

	if err := cm.cmd.Start(); err != nil {
		fmt.Println("服务器启动失败:", err)
		return 0, err
	}

	javaPID = cm.cmd.Process.Pid
	fmt.Println(javaPID)

	return javaPID, nil
}

func CmdSocket(w http.ResponseWriter, r *http.Request, cm *CommandManager) {
	// var outputLines []string
	var output string
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for cm.stdout.Scan() {
			output = cm.stdout.Text()
			fmt.Println(output) // 将命令行输出打印到控制台
			err := conn.WriteMessage(websocket.TextMessage, []byte(output))
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	//defer conn.Close()
}
