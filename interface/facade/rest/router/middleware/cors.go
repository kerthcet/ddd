package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetCors 跨域设置
func SetCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Client-Type", "X-User-ID", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
