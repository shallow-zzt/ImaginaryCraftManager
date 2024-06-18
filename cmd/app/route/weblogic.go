package route

import (
	"ImaginaryCraftManager/cmd/app/route/rutils"
	"net/http"
)

func RouteWeblogic() {
	/* ---------------- 中间件 ----------------*/
	rg := rutils.NewRouteGroup("/redirect")
	handlers := rutils.Handlers{
		"2/dashboard": func(w http.ResponseWriter, r *http.Request) {
			CheckCookie(w, r)
		},
	}
	rg.AddRoute(handlers)
}
