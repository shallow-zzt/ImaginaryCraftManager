package route

import (
	"ImaginaryCraftManager/auth/weblogin"
	"ImaginaryCraftManager/cmd/app/route/rutils"
	"fmt"
	"net/http"
)

func RouteIndex() {
	/* ---------------- 控制台显示 ----------------*/
	rg := rutils.NewRouteGroup("/")
	handlers := rutils.Handlers{
		"dashboard": func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL.Path)
			if CheckCookie(w, r) {
				return
			}
			dashboard(w, r)
		},
		"": func(w http.ResponseWriter, r *http.Request) {
			if weblogin.CheckIsLogined(r) {
				http.Redirect(w, r, "/dashboard", http.StatusFound)
				return
			}
			loginPage(w, r)
		},
	}
	rg.AddRoute(handlers)
}
func dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("调用成功")
	http.ServeFile(w, r, "static/dashboard.html")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
