package routers

import (
	healtchcheck "0xaust1n.github.com/gin-template/internal/handler/healthcheck"
	"0xaust1n.github.com/gin-template/internal/pkg/core"
)

func Register(s *core.Server) {
	healthCheckHandler := healtchcheck.NewHealthCheck(s.Redis)
	s.Router.GET("/health", healthCheckHandler.Ok)
}
