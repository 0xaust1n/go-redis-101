package healtchcheck

import (
	"0xaust1n.github.com/gin-template/internal/pkg/core"
	"github.com/gin-gonic/gin"
)

var _ IHealthCheck = (*healthCheckHandler)(nil)

type IHealthCheck interface {
	Ok(c *gin.Context)
	i()
}

type healthCheckHandler struct {
	redis *core.Redis
}

func NewHealthCheck(redis *core.Redis) IHealthCheck {
	return &healthCheckHandler{
		redis: redis,
	}
}

func (h *healthCheckHandler) i() {}
