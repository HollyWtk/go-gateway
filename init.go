package main

import (
	"gateway/config"
	"gateway/middleware"
	"gateway/nacos"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Init(router *gin.Engine) {
	nacos.InitNacosServer()
	initRouter(router)
	initServer(router)
}

func initRouter(router *gin.Engine) {
	router.Use(middleware.Request(), middleware.Cors())
}

func initServer(router *gin.Engine) {
	host := config.File.MustValue("web_server", "host", "127.0.0.1")
	port := config.File.MustValue("web_server", "port", "8088")
	s := &http.Server{
		Addr:           host + ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	log.Println(err)
}
