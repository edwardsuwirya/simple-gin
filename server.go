package main

import (
	"enigmacamp.com/simplegin/config"
	"enigmacamp.com/simplegin/delivery/api"
	"enigmacamp.com/simplegin/delivery/middleware"
	"github.com/gin-gonic/gin"
)

type AppServer interface {
	Run()
}

type productServer struct {
	router *gin.Engine
	config config.Config
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
	listenAddr := p.config.Get("productapp.api.url")
	err := p.router.Run(listenAddr)
	if err != nil {
		panic(err)
	}
}

func Server() AppServer {
	r := gin.Default()
	c := config.New()
	return &productServer{
		router: r,
		config: c,
	}
}
