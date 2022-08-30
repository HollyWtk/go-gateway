package predicate

import (
	"gateway/net"
	"github.com/gin-gonic/gin"
)

var PathPredicate = &Path{}

type Path struct {
}

func (p *Path) Predicate(predicate *net.Predicate) {
	g := predicate.Group("lspos-finance")
	g.AddPredicate("path", p.convertPath)
}

func (p *Path) convertPath(ctx *gin.Context) {
	println(ctx.Request.URL)
}
