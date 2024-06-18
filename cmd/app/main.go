package app

import (
	"ImaginaryCraftManager/auth/weblogin"
	"ImaginaryCraftManager/cmd/app/route"
	logger "ImaginaryCraftManager/log"
	"net/http"
)

const (
	// 因为我不确定需要添加到配置哪里，所以就先暂时写到这里了。
	tlsEnable   = false   // 控制是否启用tls。
	loggerLevel = "debug" // 日志等级
)

func Main() {
	// 启动日志器
	logger.NewLogger(loggerLevel)

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

	// 启动HTTP服务
	if tlsEnable {
		err := http.ListenAndServeTLS(":8080", "server.cert", "server.key", nil)
		if err != nil {
			logger.Fatalf("Main: 开启HTTP in TLS时遇到错误: %v", err)
			return
		}
	} else {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			logger.Fatalf("Main: 开启HTTP服务时遇到错误: %v", err)
			return
		}
	}

	logger.Infoln("running...(*^▽^*)")
}
