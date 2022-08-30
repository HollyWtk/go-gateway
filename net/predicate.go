package net

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"sync"
)

type HandlerFunc func(ctx *gin.Context)
type MiddlewareFunc func(handlerFunc HandlerFunc) HandlerFunc

type group struct {
	mutex         sync.RWMutex
	prefix        string
	handlerMap    map[string]HandlerFunc
	middlewareMap map[string][]MiddlewareFunc
	middlewares   []MiddlewareFunc
}

func (g *group) AddPredicate(name string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.handlerMap[name] = handlerFunc
	g.middlewareMap[name] = middlewares
}

func (g *group) Use(middlewares ...MiddlewareFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}
func (g *group) exec(name string, ctx *gin.Context) {
	h, ok := g.handlerMap[name]
	if !ok {
		log.Println("路由未定义")
	} else {
		for i := 0; i < len(g.middlewares); i++ {
			h = g.middlewares[i](h)
		}
		mm, ok := g.middlewareMap[name]
		if ok {
			for i := 0; i < len(mm); i++ {
				h = mm[i](h)
			}
		}
		h(ctx)
	}
}

func (p *Predicate) Group(prefix string) *group {
	g := &group{
		prefix:        prefix,
		handlerMap:    make(map[string]HandlerFunc),
		middlewareMap: make(map[string][]MiddlewareFunc),
	}
	p.group = append(p.group, g)
	return g
}

type Predicate struct {
	group []*group
}

func NewPredicate() *Predicate {
	return &Predicate{}
}

func (p *Predicate) Run(ctx *gin.Context) {
	realPath := ctx.Request.URL.Path[1:]
	var serviceName string
	index := strings.Index(realPath, "/")
	if index > 0 {
		serviceName = realPath[:index]
	} else {
		serviceName = realPath
	}
	for _, g := range p.group {
		if g.prefix == serviceName {
			g.exec(serviceName, ctx)
		}
	}
}
