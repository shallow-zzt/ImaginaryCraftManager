package route

import (
	"ImaginaryCraftManager/cmd/app/route/rutils"
	"ImaginaryCraftManager/generic/serverCmd"
	"fmt"
	"net/http"
)

var manager *serverCmd.CommandManager
var serverRunning bool

func RouteControl() {
	/* ---------------- 服务器控制url接口 ----------------*/
	rg := rutils.NewRouteGroup("control")
	handlers := rutils.Handlers{
		"servercmd/start": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			//启动服务器
			if !serverRunning {
				var err error
				manager, err = serverCmd.NewCmdManager("fabric-server")
				if err != nil {
					return
				}
				serverRunning = true
				startCmd(w, r, manager)

			} else {
				fmt.Println("进程已启动")
			}
		},
		"servercmd/restart": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			//重启服务器
			if serverRunning {
				stopCmd(w, r, manager)
				var err error
				manager, err = serverCmd.NewCmdManager("fabric-server")
				if err != nil {
					return
				}
				startCmd(w, r, manager)
			} else {
				fmt.Println("进程未启动")
			}

		},
		"servercmd/stop": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			//关闭服务器
			if serverRunning {
				serverRunning = false
				stopCmd(w, r, manager)
			} else {
				fmt.Println("进程未启动")
			}

		},
	}
	rg.AddRoute(handlers)
}

func startCmd(w http.ResponseWriter, r *http.Request, cm *serverCmd.CommandManager) {
	javaPid, err := serverCmd.CmdRecording(w, r, cm)
	fmt.Println(javaPid)
	if err != nil {
		return
	}
}

func showCmd(w http.ResponseWriter, r *http.Request, cm *serverCmd.CommandManager) {
	serverCmd.CmdSocket(w, r, cm)
}

func stopCmd(w http.ResponseWriter, r *http.Request, cm *serverCmd.CommandManager) {

	err := serverCmd.CloseProcessAndPipe(cm)
	if err != nil {
		return
	}
}
