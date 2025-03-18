package example

import (
	"0xaust1n.github.com/gin-template/internal/pkg/core"
	"github.com/gin-gonic/gin"
)

var _ IHeadler = (*handler)(nil)

type IHeadler interface {
	Ok(c *gin.Context)
	i()
}

type handler struct {
	redis *core.Redis
}

func NewHealthCheck(redis *core.Redis) IHeadler {
	return &handler{
		redis: redis,
	}
}

func (h *handler) i() {}
