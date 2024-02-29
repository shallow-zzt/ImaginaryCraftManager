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

func main() {
	// 设置路由
	var err error
	if err != nil {
		fmt.Println("cmd管道创建失败:", err)
		return
	}

	http.HandleFunc("/api/mods", func(w http.ResponseWriter, r *http.Request) {
		ShowMods(w, r)
	})
	http.HandleFunc("/control/servercmd/start", func(w http.ResponseWriter, r *http.Request) {
		manager, err = serverCmd.NewCmdManager("fabric-server")
		StartCmd(w, r, manager)
	})
	http.HandleFunc("/control/servercmd/stop", func(w http.ResponseWriter, r *http.Request) {
		if manager == nil {
			fmt.Println("Manager 未初始化")
			return
		}
		StopCmd(w, r, manager)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
