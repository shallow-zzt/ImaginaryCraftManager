package main

import (
	"ImaginaryCraftManager/auth/authStructs"
	"ImaginaryCraftManager/auth/weblogin"
	"ImaginaryCraftManager/generic/fileManage"
	"ImaginaryCraftManager/generic/serverCmd"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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
	/* ---------------- web访问控制 ----------------*/
	http.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		//初始登陆界面
		LoginFunc(w, r)
	})

	/* ---------------- api查询url接口 ----------------*/
	http.HandleFunc("/api/mods", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
		//mod列表展示
		ShowMods(w, r)
	})
	http.HandleFunc("/api/mods/configs", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
		//mod配置列表展示
		ShowModsConfigs(w, r)
	})
	http.HandleFunc("/api/server/setting", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
		//服务器设置查看
	})
	http.HandleFunc("/api/server/status", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
		//服务器状态查看
	})

	/* ---------------- 设置修改url接口 ----------------*/
	http.HandleFunc("/setting/modify/servercmd/gamerule", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
		//服务器游戏设置修改
	})
	http.HandleFunc("/setting/modify/servercmd/running", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
		//服务器启动设置
	})

	/* ---------------- 文件控制url接口 ----------------*/
	http.HandleFunc("/file/mods/upload", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
		//服务器mod上传
	})
	http.HandleFunc("/file/mods/download", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
		//服务器mod下载
	})
	http.HandleFunc("/file/mods/delete", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
		//服务器mod删除
	})
	http.HandleFunc("/file/mods/config/upload", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
		//服务器mod配置上传
	})
	http.HandleFunc("/file/mods/config/download", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
		//服务器mod配置下载
	})
	http.HandleFunc("/file/mods/config/delete", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
		//服务器mod配置删除
	})

	/* ---------------- 服务器控制url接口 ----------------*/
	http.HandleFunc("/control/servercmd/start", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
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
		if RedirectHandler(w, r) {
			return
		}
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
		if RedirectHandler(w, r) {
			return
		}
		//关闭服务器
		if serverRunning {
			serverRunning = false
			StopCmd(w, r, manager)
		} else {
			fmt.Println("进程未启动")
		}

	})

	/* ---------------- 控制台显示 ----------------*/
	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		if RedirectHandler(w, r) {
			return
		}
		Dashboard(w, r)
	})

	/* ---------------- 中间件 ----------------*/
	http.HandleFunc("/redirect/2/dashboard", func(w http.ResponseWriter, r *http.Request) {
		RedirectHandler(w, r)
	})

	/* ---------------- 静态资源路径 ----------------*/
	http.Handle("/", http.FileServer(http.Dir("static")))

	// 启动Web服务器
	http.ListenAndServe(":8080", nil)
	fmt.Println("running……")
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) bool {
	if !weblogin.CheckIsLogined(r) {
		http.Redirect(w, r, "/", http.StatusFound)
		return true
	}
	return false
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("调用成功")
	http.ServeFile(w, r, "static/dashboard.html")
}

// 登录处理函数
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	var user authStructs.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "非法请求体", http.StatusBadRequest)
		return
	}

	if weblogin.CheckLogin(user.Username, user.Password) {
		fmt.Println("登陆成功")
		expiration := time.Now().Add(24 * time.Hour)
		cookie := &http.Cookie{
			Name:    "session",
			Value:   user.Username,
			Expires: expiration,
			Path:    "/",
		}
		http.SetCookie(w, cookie)
		//http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	}
	http.Error(w, "账号或者密码错误", http.StatusUnauthorized)
	//http.Redirect(w, r, "/", http.StatusFound)

}

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

func ShowModsConfigs(w http.ResponseWriter, r *http.Request) {
	type ModPath struct {
		Configs []string `json:"configs"`
	}
	configPath := "fabric-server/config"
	var configLists []string

	fileNames, err := fileManage.GetAllFileNames(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	configLists = append(configLists, fileNames...)
	response := ModPath{Configs: configLists}

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
