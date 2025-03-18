package healtchcheck

import (
	"net/http"

	"0xaust1n.github.com/gin-template/internal/interfaces"
	"github.com/gin-gonic/gin"
)

func (h *healthCheckHandler) Ok(c *gin.Context) {
	message := "Server Alive"

	h.redis.Set("healthcheck", "ok", 0)

	c.JSON(http.StatusOK, interfaces.BaseResponse{
		Code:    http.StatusOK,
		Message: &message,
	})
}
