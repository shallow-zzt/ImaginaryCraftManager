package main

import (
	"ImaginaryCraftManager/generic/fileManage"
	"ImaginaryCraftManager/generic/serverCmd"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

var cmd *exec.Cmd = exec.Command("cmd.exe", "/c", "start.bat")

func main() {
	// 设置路由
	http.HandleFunc("/api/mods", ShowMods)

	http.HandleFunc("/control/start", StartCmd)
	http.HandleFunc("/control/stop", StopCmd)

	http.Handle("/", http.FileServer(http.Dir("static")))

	// 启动Web服务器
	http.ListenAndServe(":8080", nil)
	fmt.Println("running……")
}

// showMode 处理从前端发送的消息
func ShowMods(w http.ResponseWriter, r *http.Request) {
	// 从请求中解析JSON数据
	type ModPath struct {
		Mods []string `json:"mods"`
	}
	modPath := "fabric-server/mods"
	var modLists []string

	fileNames, err := fileManage.GetAllFileNames(modPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	modLists = append(modLists, fileNames...)
	response := ModPath{Mods: modLists}

	// 将响应数据转换为JSON格式并写回到客户端
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func StartCmd(w http.ResponseWriter, r *http.Request) {
	cmd.Dir = "fabric-server"
	javaPid, err := serverCmd.CmdRecording(cmd)
	fmt.Println(javaPid)
	if err != nil {
		return
	}
}

func StopCmd(w http.ResponseWriter, r *http.Request) {
	cmd.Dir = "fabric-server"
	//暴力关闭Java
	//因为我也没想到更好的办法
	//我不清楚这样做，会不会把所有需要java运行的程序都关了 ^_^
	exec.Command("taskkill", "/f", "/im", "java.exe").Run()
	err := serverCmd.CloseProcessAndPipe(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
