package net

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type server struct {
	addr       string
	predicate  *Predicate
	needSecret bool
}

func (s *server) Predicate(predicate *Predicate) {
	s.predicate = predicate
}

func NewServer(addr string) *server {
	return &server{
		addr: addr,
	}
}
func (s *server) NeedSecret(needSecret bool) {
	s.needSecret = needSecret
}

func (s *server) Start(router *gin.Engine) {
	server := &http.Server{
		Addr:           s.addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
