package http

import (
	"github.com/go-jones/restful-api-demo/app/host"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/julienschmidt/httprouter"
)

// host 模块的HTTP API 暴露
var api = &handler{}

type handler struct {
	host host.Service
	log  logger.Logger
}

// 初始化的时候，依赖外部Host Sevice的实例对象
func (h *handler) Init() {
	h.log = zap.L().Named("HOST API")
}
func (h *handler) Registry(r *httprouter.Router) {
	r.POST("/hosts", h.CreateHost)
	r.GET("/hosts", h.QueryHost)
}
