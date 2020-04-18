package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"ddd/infrastructure/util/migrate"
	"ddd/interface/facade/rest/router"
	"ddd/interface/facade/rest/router/middleware"
	"ddd/interface/facade/rpc/client"
	"ddd/interface/facade/rpc/server"
)

func main() {
	migrate.Run()

	g := gin.New()

	middlewares := []gin.HandlerFunc{
		middleware.SetCors(),
	}

	router.Load(
		g,
		middlewares...,
	)

	log.Println("Grpc listen on port: 9090")
	go server.Run()

	client.Send()
	log.Printf("Server listen on port: 8080")
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("%+v", err)
	}
}
