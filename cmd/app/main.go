package app

import (
	"ImaginaryCraftManager/auth/weblogin"
	"ImaginaryCraftManager/cmd/app/route"
	logger "ImaginaryCraftManager/log"
	"fmt"
	"net/http"
)

// 因为我不确定需要添加到配置哪里，所以就先暂时写到这里了。控制是否启用tls。
const tlsEnable = false

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

	if tlsEnable {
		err := http.ListenAndServeTLS(":8080", "server.cert", "server.key", nil)
		if err != nil {
			logger.Fatalf("Main: 开启TLS时遇到错误: %v", err)
			return
		}
	} else {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			logger.Fatalf("Main: 开启HTTP服务时遇到错误: %v", err)
			return
		}
	}

	fmt.Println("running...(*^▽^*)")
}
