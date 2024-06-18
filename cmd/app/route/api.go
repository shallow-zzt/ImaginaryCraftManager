package route

import (
	"ImaginaryCraftManager/auth/weblogin"
	"ImaginaryCraftManager/cmd/app/route/rutils"
	"ImaginaryCraftManager/generic/fileManage"
	"ImaginaryCraftManager/generic/serverConfig"
	"ImaginaryCraftManager/jsonStructs/responseStructs/pathStructs"
	logger "ImaginaryCraftManager/log"
	"encoding/json"
	"fmt"
	"net/http"
)

func RouteApi() {
	/* ---------------- api查询url接口 ----------------*/
	rg := rutils.NewRouteGroup("/api")
	handlers := rutils.Handlers{
		"mods": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			//mod列表展示
			showMods(w, r)
		},
		"mods/configs": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			//mod配置列表展示
			showModsConfigs(w, r)
		},
		"server/setting": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			showServerConfig(w, r)
			//服务器设置查看
		},
		"server/status": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			//服务器状态查看
		},
	}
	rg.AddRoute(handlers)
}

func CheckCookie(w http.ResponseWriter, r *http.Request) bool {
	if !weblogin.CheckIsLogined(r) {
		http.Redirect(w, r, "/", http.StatusFound)
		return true
	}
	return false
}

func showServerConfig(w http.ResponseWriter, r *http.Request) {
	serverConfigList, err := serverConfig.ReadServerConfig("fabric-server/server.properties")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	response := serverConfig.WriteServerConfig2Json(serverConfigList, "fabric-server")
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func showMods(w http.ResponseWriter, r *http.Request) {
	modPath := "fabric-server/mods"
	var modLists []string

	fileNames, err := fileManage.GetAllFileNames(modPath)
	if err != nil {
		logger.Fatalln(err)
		return
	}

	modLists = append(modLists, fileNames...)

	response := pathStructs.ModPath{Mods: modLists, ModNums: len(modLists)}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func showModsConfigs(w http.ResponseWriter, r *http.Request) {
	configPath := "fabric-server/config"
	var configLists []string

	fileNames, err := fileManage.GetAllFileNames(configPath)
	if err != nil {
		logger.Fatalln(err)
		return
	}

	configLists = append(configLists, fileNames...)

	response := pathStructs.ModConfigPath{Configs: configLists, ConfigNums: len(configLists)}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
