package middleware

import (
	"gateway/net"
	"github.com/gin-gonic/gin"
)

func DoPredicate(predicate *net.Predicate) gin.HandlerFunc {
	return func(c *gin.Context) {
		predicate.Run(c)
		c.Next()
	}
}
