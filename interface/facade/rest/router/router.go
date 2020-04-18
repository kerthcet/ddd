package router

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"ddd/interface/facade/rest/handler"
)

// Load 加载router
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(mw...)

	pprof.Register(g)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Incorrect route")
	})

	ping := g.Group("/ping")
	{
		ping.GET("", handler.Ping)
	}

	api := g.Group("/api")
	{
		api.POST("leave", handler.Leave)
	}
	return g
}
