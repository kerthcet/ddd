package handler

import (
	"github.com/gin-gonic/gin"
)

// Ping health check
func Ping(c *gin.Context) {
	SendResponse(c, nil, "pong")
}
