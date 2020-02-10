package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"ddd/infrastructure/config"
	// "ddd/infrastructure/util/driver"
	"ddd/interface/facade/rest/router"
	"ddd/interface/facade/rest/router/middleware"
)

func main() {
	// driver.DB.AutoMigrate(&po.Leave{})

	g := gin.New()
	log.Println(config.ENV.GinMode)

	middlewares := []gin.HandlerFunc{
		middleware.SetCors(),
	}

	router.Load(
		g,
		middlewares...,
	)

	log.Printf("Start to listening port: 8080")
	if err := g.Run(":8080"); err != nil {
		log.Print(err)
		panic(err)
	}
}
