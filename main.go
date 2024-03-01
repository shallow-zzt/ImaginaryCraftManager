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

	/* ---------------- api查询url接口 ----------------*/
	http.HandleFunc("/api/mods", func(w http.ResponseWriter, r *http.Request) {
		//mod列表展示
		ShowMods(w, r)
	})
	http.HandleFunc("/api/mods/configs", func(w http.ResponseWriter, r *http.Request) {
		//mod配置列表展示
	})
	http.HandleFunc("/api/server/setting", func(w http.ResponseWriter, r *http.Request) {
		//服务器设置查看
	})
	http.HandleFunc("/api/server/status", func(w http.ResponseWriter, r *http.Request) {
		//服务器状态查看
	})

	/* ---------------- 设置修改url接口 ----------------*/
	http.HandleFunc("/setting/modify/servercmd/gamerule", func(w http.ResponseWriter, r *http.Request) {
		//服务器游戏设置修改
	})
	http.HandleFunc("/setting/modify/servercmd/running", func(w http.ResponseWriter, r *http.Request) {
		//服务器启动设置
	})

	/* ---------------- 文件控制url接口 ----------------*/
	http.HandleFunc("/file/mods/upload", func(w http.ResponseWriter, r *http.Request) {
		//服务器mod上传
	})
	http.HandleFunc("/file/mods/download", func(w http.ResponseWriter, r *http.Request) {
		//服务器mod下载
	})
	http.HandleFunc("/file/mods/delete", func(w http.ResponseWriter, r *http.Request) {
		//服务器mod删除
	})
	http.HandleFunc("/file/mods/config/upload", func(w http.ResponseWriter, r *http.Request) {
		//服务器mod配置上传
	})
	http.HandleFunc("/file/mods/config/download", func(w http.ResponseWriter, r *http.Request) {
		//服务器mod配置下载
	})
	http.HandleFunc("/file/mods/config/delete", func(w http.ResponseWriter, r *http.Request) {
		//服务器mod配置删除
	})

	/* ---------------- 服务器控制url接口 ----------------*/
	http.HandleFunc("/control/servercmd/start", func(w http.ResponseWriter, r *http.Request) {
		//启动服务器
		if !serverRunning {
			manager, err = serverCmd.NewCmdManager("fabric-server")
			serverRunning = true
			StartCmd(w, r, manager)

		} else {
			fmt.Println("进程已启动")
		}

	})
	http.HandleFunc("/control/servercmd/restart", func(w http.ResponseWriter, r *http.Request) {
		//重启服务器
		if serverRunning {
			StopCmd(w, r, manager)
			manager, err = serverCmd.NewCmdManager("fabric-server")
			StartCmd(w, r, manager)
		} else {
			fmt.Println("进程未启动")
		}

	})
	http.HandleFunc("/control/servercmd/stop", func(w http.ResponseWriter, r *http.Request) {
		//关闭服务器
		if serverRunning {
			serverRunning = false
			StopCmd(w, r, manager)
		} else {
			fmt.Println("进程未启动")
		}

	})

	/* ---------------- 静态资源路径 ----------------*/
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
