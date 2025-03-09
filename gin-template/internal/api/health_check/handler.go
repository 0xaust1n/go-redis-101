package healthCheck

import (
	"net/http"

	"0xaust1n.github.com/gin-template/internal/interfaces"
	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context) {
	message := "Server Alive"
	c.JSON(http.StatusOK, interfaces.BaseResponse{
		Code:    http.StatusOK,
		Message: &message,
	})
}
