package route

import (
	"ImaginaryCraftManager/cmd/app/route/rutils"
	"net/http"
	"path/filepath"
)

func RouteStatic() {
	/* ---------------- 静态资源路径 ----------------*/
	routeJs()
	routeCss()
}

func routeJs() {
	rg := rutils.NewRouteGroup("/js")
	handlers := rutils.Handlers{
		"backendFunc.js": serveStaticFile("backendFunc.js", "js"),
		"login.js":       serveStaticFile("login.js", "js"),
		"dashboard.js":   serveStaticFile("dashboard.js", "js"),
		"websocket.js":   serveStaticFile("websocket.js", "js"),
	}
	rg.AddRoute(handlers)
}

func routeCss() {
	rg := rutils.NewRouteGroup("/css")
	handlers := rutils.Handlers{
		"dashboard.css": serveStaticFile("dashboard.css", "css"),
	}
	rg.AddRoute(handlers)
}

func serveStaticFile(filename string, fileExt string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if fileExt == "js" {
			w.Header().Set("Content-Type", "application/javascript")
			http.ServeFile(w, r, filepath.Join("static/js", filename))
		} else if fileExt == "css" {
			w.Header().Set("Content-Type", "text/css")
			http.ServeFile(w, r, filepath.Join("static/css", filename))
		}
	}
}
