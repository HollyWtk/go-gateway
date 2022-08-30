package main

import (
	"gateway/config"
	"gateway/middleware"
	"gateway/nacos"
	"gateway/net"
	"gateway/net/predicate"
	"github.com/gin-gonic/gin"
	"strconv"
)

var DefaultPredicate = net.NewPredicate()

func Init(router *gin.Engine) {
	nacos.InitNacosServer()
	server := net.NewServer(config.WebServerConfig.Host + ":" + strconv.FormatInt(config.WebServerConfig.Port, 10))
	initPredicate()
	server.Predicate(DefaultPredicate)
	router.Use(middleware.DoPredicate(DefaultPredicate), middleware.Request(), middleware.Cors())
	server.Start(router)
}

func initPredicate() {
	predicate.PathPredicate.Predicate(DefaultPredicate)
}
