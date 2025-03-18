package routers

import (
	"0xaust1n.github.com/gin-template/internal/handler/example"
	"0xaust1n.github.com/gin-template/internal/pkg/core"
)

func Register(s *core.Server) {
	exampleHandler := example.NewHealthCheck(s.Redis)
	s.Router.GET("/health", exampleHandler.Ok)
}
