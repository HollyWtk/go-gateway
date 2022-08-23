package middleware

import (
	"github.com/gin-gonic/gin"
)

func Request() gin.HandlerFunc {
	return func(c *gin.Context) {
		//nacos.GetConfig()
		println(c.Request.URL.String())
	}
}
