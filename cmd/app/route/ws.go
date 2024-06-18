package route

import (
	"ImaginaryCraftManager/cmd/app/route/rutils"
	"fmt"
	"net/http"
)

func RouteWs() {
	/* ---------------- websocket ----------------*/
	rg := rutils.NewRouteGroup("/ws")
	handlers := rutils.Handlers{
		"servercmd/status": func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL.Path)
			if CheckCookie(w, r) {
				return
			}
			if serverRunning {
				showCmd(w, r, manager)
			} else {
				fmt.Println("进程未启动")
			}

		},
	}
	rg.AddRoute(handlers)
}
