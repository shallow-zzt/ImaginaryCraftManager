package route

import (
	"ImaginaryCraftManager/cmd/app/route/rutils"
	"ImaginaryCraftManager/generic/serverCmd"
	"ImaginaryCraftManager/generic/serverConfig"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func RouteSetting() {
	/* ---------------- 设置修改url接口 ----------------*/
	rg := rutils.NewRouteGroup("/setting/modify/servercmd")
	handlers := rutils.Handlers{
		"configs": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			setServerConfigs(w, r)
			//服务器游戏设置修改
		},
		"running": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			//服务器启动设置
		},
	}
	rg.AddRoute(handlers)
}
func setServerConfigs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	var settingStruct map[string]string
	err := json.NewDecoder(r.Body).Decode(&settingStruct)
	if err != nil {
		http.Error(w, "非法请求体", http.StatusBadRequest)
		return
	}

	var memorysize int
	for key, value := range settingStruct {
		if key == "server-memory" {
			memorysize, err = strconv.Atoi(value)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = serverCmd.SetCmdParameter("fabric-server", memorysize)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			err = serverConfig.WriteProperty("fabric-server/server.properties", key, value)
			if err != nil {
				return
			}
		}

	}
	w.Header().Set("Content-Type", "application/json")

}
