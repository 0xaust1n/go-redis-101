package routers

import (
	healtchcheck "0xaust1n.github.com/gin-template/internal/api/healthcheck"
	"0xaust1n.github.com/gin-template/internal/pkg/core"
)

func RegisterRouters(s *core.Server) {
	healthCheckHandler := healtchcheck.NewHealthCheck(s.Redis)
	s.Router.GET("/health", healthCheckHandler.Ok)
}
