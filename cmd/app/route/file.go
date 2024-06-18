package route

import (
	"ImaginaryCraftManager/cmd/app/route/rutils"
	"net/http"
)

func RouteFile() {
	/* ---------------- 文件控制url接口 ----------------*/
	rg := rutils.NewRouteGroup("/file/mods")
	handlers := rutils.Handlers{
		"upload": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			//服务器mod上传
		},
		"download": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			//服务器mod下载
		},
		"delete": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			//服务器mod删除
		},
		"config/upload": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			//服务器mod配置上传
		},
		"config/download": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			//服务器mod配置下载
		},
		"config/delete": func(w http.ResponseWriter, r *http.Request) {
			if CheckCookie(w, r) {
				return
			}
			//服务器mod配置删除
		},
	}
	rg.AddRoute(handlers)
}
