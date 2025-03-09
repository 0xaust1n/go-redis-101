package routers

import (
	healthCheck "0xaust1n.github.com/gin-template/internal/api/health_check"
	"0xaust1n.github.com/gin-template/internal/pkg/core"
)

func RegisterRouters(r *core.Router) {
	r.GET("/", healthCheck.Ok)
}
