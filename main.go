package main

import (
	"ImaginaryCraftManager/generic/fileManage"
	"ImaginaryCraftManager/generic/serverCmd"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var manager *serverCmd.CommandManager
var serverRunning bool

func main() {
	// 设置路由
	var err error
	serverRunning = false
	if err != nil {
		return
	}

	http.HandleFunc("/api/mods", func(w http.ResponseWriter, r *http.Request) {
		ShowMods(w, r)
	})
	http.HandleFunc("/control/servercmd/start", func(w http.ResponseWriter, r *http.Request) {
		if !serverRunning {
			manager, err = serverCmd.NewCmdManager("fabric-server")
			serverRunning = true
			StartCmd(w, r, manager)

		} else {
			fmt.Println("进程已启动")
		}

	})
	http.HandleFunc("/control/servercmd/stop", func(w http.ResponseWriter, r *http.Request) {
		if serverRunning {
			// if manager == nil {
			// 	fmt.Println("Manager 未初始化")
			// 	return
			// }
			serverRunning = false
			StopCmd(w, r, manager)
		} else {
			fmt.Println("进程未启动")
		}

	})

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

func StartCmd(w http.ResponseWriter, r *http.Request, cm *serverCmd.CommandManager) {
	javaPid, err := serverCmd.CmdRecording(cm)
	fmt.Println(javaPid)
	if err != nil {
		return
	}
}

func StopCmd(w http.ResponseWriter, r *http.Request, cm *serverCmd.CommandManager) {

	err := serverCmd.CloseProcessAndPipe(cm)
	if err != nil {
		return
	}
}
