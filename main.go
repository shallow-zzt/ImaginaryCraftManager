package main

import (
	"ImaginaryCraftManager/auth/weblogin"
	"ImaginaryCraftManager/generic/fileManage"
	"ImaginaryCraftManager/generic/serverCmd"
	"ImaginaryCraftManager/generic/serverConfig"
	"ImaginaryCraftManager/jsonStructs/requestStructs/authStructs"
	"ImaginaryCraftManager/jsonStructs/responseStructs/pathStructs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
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
	//taskKillProtect.InterruptFunction(manager)
	weblogin.LoadUsers("authorities.ini")
	/* ---------------- web访问控制 ----------------*/
	http.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		//初始登陆界面
		LoginFunc(w, r)
	})
	http.HandleFunc("/auth/logout", func(w http.ResponseWriter, r *http.Request) {
		//登出账号
		if CheckCookie(w, r) {
			return
		}
		Logoutfunc(w, r)
	})
	http.HandleFunc("/auth/logout/refresh", func(w http.ResponseWriter, r *http.Request) {
		//刷新账号信息并登出
		if CheckCookie(w, r) {
			return
		}
		RefreshLogin(w, r)
	})

	/* ---------------- api查询url接口 ----------------*/
	http.HandleFunc("/api/mods", func(w http.ResponseWriter, r *http.Request) {
		if CheckCookie(w, r) {
			return
		}
		//mod列表展示
		ShowMods(w, r)
	})
	http.HandleFunc("/api/mods/configs", func(w http.ResponseWriter, r *http.Request) {
		if CheckCookie(w, r) {
			return
		}
		//mod配置列表展示
		ShowModsConfigs(w, r)
	})
	http.HandleFunc("/api/server/setting", func(w http.ResponseWriter, r *http.Request) {
		if CheckCookie(w, r) {
			return
		}
		ShowServerConfig(w, r)
		//服务器设置查看
	})
	http.HandleFunc("/api/server/status", func(w http.ResponseWriter, r *http.Request) {
		if CheckCookie(w, r) {
			return
		}
		//服务器状态查看
	})

	/* ---------------- 设置修改url接口 ----------------*/
	http.HandleFunc("/setting/modify/servercmd/gamerule", func(w http.ResponseWriter, r *http.Request) {
		if CheckCookie(w, r) {
			return
		}
		//服务器游戏设置修改
	})
	http.HandleFunc("/setting/modify/servercmd/running", func(w http.ResponseWriter, r *http.Request) {
		if CheckCookie(w, r) {
			return
		}
		//服务器启动设置
	})

	/* ---------------- 文件控制url接口 ----------------*/
	http.HandleFunc("/file/mods/upload", func(w http.ResponseWriter, r *http.Request) {
		if CheckCookie(w, r) {
			return
		}
		//服务器mod上传
	})
	http.HandleFunc("/file/mods/download", func(w http.ResponseWriter, r *http.Request) {
		if CheckCookie(w, r) {
			return
		}
		//服务器mod下载
	})
	http.HandleFunc("/file/mods/delete", func(w http.ResponseWriter, r *http.Request) {
		if CheckCookie(w, r) {
			return
		}
		//服务器mod删除
	})
	http.HandleFunc("/file/mods/config/upload", func(w http.ResponseWriter, r *http.Request) {
		if CheckCookie(w, r) {
			return
		}
		//服务器mod配置上传
	})
	http.HandleFunc("/file/mods/config/download", func(w http.ResponseWriter, r *http.Request) {
		if CheckCookie(w, r) {
			return
		}
		//服务器mod配置下载
	})
	http.HandleFunc("/file/mods/config/delete", func(w http.ResponseWriter, r *http.Request) {
		if CheckCookie(w, r) {
			return
		}
		//服务器mod配置删除
	})

	/* ---------------- 服务器控制url接口 ----------------*/
	http.HandleFunc("/control/servercmd/start", func(w http.ResponseWriter, r *http.Request) {
		if CheckCookie(w, r) {
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
		if CheckCookie(w, r) {
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
		if CheckCookie(w, r) {
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

	/* ---------------- websocket ----------------*/
	http.HandleFunc("/ws/servercmd/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if CheckCookie(w, r) {
			return
		}
		if serverRunning {
			ShowCmd(w, r, manager)
		} else {
			fmt.Println("进程未启动")
		}

	})

	/* ---------------- 控制台显示 ----------------*/
	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if CheckCookie(w, r) {
			return
		}
		Dashboard(w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if weblogin.CheckIsLogined(r) {
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		}
		LoginPage(w, r)
	})

	/* ---------------- 中间件 ----------------*/
	http.HandleFunc("/redirect/2/dashboard", func(w http.ResponseWriter, r *http.Request) {
		CheckCookie(w, r)
	})

	/* ---------------- 静态资源路径 ----------------*/
	http.HandleFunc("/js/backendFunc.js", ServeStaticFile("backendFunc.js", "js"))
	http.HandleFunc("/js/login.js", ServeStaticFile("login.js", "js"))
	http.HandleFunc("/js/dashboard.js", ServeStaticFile("dashboard.js", "js"))
	http.HandleFunc("/js/websocket.js", ServeStaticFile("websocket.js", "js"))
	http.HandleFunc("/css/dashboard.css", ServeStaticFile("dashboard.css", "css"))

	// 启动Web服务器
	http.ListenAndServe(":8080", nil)
	fmt.Println("running……")
}

func ServeStaticFile(filename string, fileext string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if fileext == "js" {
			w.Header().Set("Content-Type", "application/javascript")
			http.ServeFile(w, r, filepath.Join("static/js", filename))
		} else if fileext == "css" {
			w.Header().Set("Content-Type", "text/css")
			http.ServeFile(w, r, filepath.Join("static/css", filename))
		}
	}
}

func CheckCookie(w http.ResponseWriter, r *http.Request) bool {
	if !weblogin.CheckIsLogined(r) {
		http.Redirect(w, r, "/", http.StatusFound)
		return true
	}
	return false
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("调用成功")
	http.ServeFile(w, r, "static/dashboard.html")
}

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
		return
	}
	http.Error(w, "账号或者密码错误", http.StatusUnauthorized)
}

func RefreshLogin(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:    "session",
		Value:   "",
		Expires: time.Unix(1, 0),
		MaxAge:  -1,
		Path:    "/",
	}
	http.SetCookie(w, cookie)
	weblogin.RefreshUsers("authorities.ini")
	weblogin.LoadUsers("authorities.ini")
}

func Logoutfunc(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:    "session",
		Value:   "",
		Expires: time.Unix(1, 0),
		MaxAge:  -1,
		Path:    "/",
	}
	http.SetCookie(w, cookie)
}

func ShowServerConfig(w http.ResponseWriter, r *http.Request) {
	serverConfigList, err := serverConfig.ReadServerConfig("fabric-server/server.properties")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	response := serverConfig.WriteServerConfig2Json(serverConfigList, "fabric-server")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ShowMods(w http.ResponseWriter, r *http.Request) {
	modPath := "fabric-server/mods"
	var modLists []string

	fileNames, err := fileManage.GetAllFileNames(modPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	modLists = append(modLists, fileNames...)

	response := pathStructs.ModPath{Mods: modLists}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ShowModsConfigs(w http.ResponseWriter, r *http.Request) {
	configPath := "fabric-server/config"
	var configLists []string

	fileNames, err := fileManage.GetAllFileNames(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	configLists = append(configLists, fileNames...)

	response := pathStructs.ModConfigPath{Configs: configLists}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func StartCmd(w http.ResponseWriter, r *http.Request, cm *serverCmd.CommandManager) {
	javaPid, err := serverCmd.CmdRecording(w, r, cm)
	fmt.Println(javaPid)
	if err != nil {
		return
	}
}

func ShowCmd(w http.ResponseWriter, r *http.Request, cm *serverCmd.CommandManager) {
	serverCmd.CmdSocket(w, r, cm)
}

func StopCmd(w http.ResponseWriter, r *http.Request, cm *serverCmd.CommandManager) {

	err := serverCmd.CloseProcessAndPipe(cm)
	if err != nil {
		return
	}
}
