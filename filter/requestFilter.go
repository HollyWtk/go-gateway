package filter

import "github.com/gin-gonic/gin"

var RequestFilter = &Filter{}

type Filter struct {
}

func (f *Filter) FilterHandler(ctx *gin.Context) {
	println(ctx.Request.URL)
}
