package rutils

import (
	"net/http"
	"strings"
)

type Handlers map[string]func(http.ResponseWriter, *http.Request)

// NewRouteGroup 创建一个新的路由组
func NewRouteGroup(urlPrefix string) *Group {
	if !strings.HasPrefix(urlPrefix, "/") {
		urlPrefix = "/" + urlPrefix
	}
	if !strings.HasSuffix(urlPrefix, "/") {
		urlPrefix = urlPrefix + "/"
	}
	return &Group{urlPrefix: urlPrefix}
}

type Group struct {
	urlPrefix string
}

// AddRoute 批量增加路由
func (c *Group) AddRoute(handlers Handlers) {
	for pattern, handler := range handlers {
		if strings.HasPrefix(pattern, "/") {
			pattern = pattern[1:]
		}
		http.HandleFunc(c.urlPrefix+pattern, handler)
	}
}
