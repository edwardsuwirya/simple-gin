package main

import (
	"enigmacamp.com/simplegin/delivery/api"
	"enigmacamp.com/simplegin/delivery/middleware"
	"github.com/gin-gonic/gin"
)

type AppServer interface {
	Run()
}

type productServer struct {
	router *gin.Engine
}

func (p *productServer) handlers() {
	p.router.Use(middleware.DummyMiddleware)
	p.router.Use(middleware.TokenAuthMiddleware())
	p.router.Use(middleware.ErrorMiddleware())
	p.v1()
}

func (p *productServer) v1() {
	productApiGroup := p.router.Group("/product")
	api.NewProductApi(productApiGroup)

	pingApiGroup := p.router.Group("/ping")
	api.NewPingApi(pingApiGroup)

}
func (p *productServer) Run() {
	p.handlers()
	err := p.router.Run("localhost:3000")
	if err != nil {
		panic(err)
	}
}

func Server() AppServer {
	r := gin.Default()
	return &productServer{
		router: r,
	}
}
