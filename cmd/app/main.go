package app

import (
	"ImaginaryCraftManager/auth/weblogin"
	"ImaginaryCraftManager/cmd/app/route"
	"fmt"
	"net/http"
)

func Main() {
	// 注册路由
	route.RouteApi()
	route.RouteAuth()
	route.RouteControl()
	route.RouteIndex()
	route.RouteFile()
	route.RouteSetting()
	route.RouteStatic()
	route.RouteWeblogic()
	route.RouteWs()

	weblogin.LoadUsers("authorities.ini")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
	fmt.Println("running...(*^▽^*)")
}
