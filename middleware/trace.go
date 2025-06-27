package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func Trace() func(c *gin.Context) {
	return func(c *gin.Context) {
		rid := uuid.NewV4().String()
		c.Set("requestId", rid)
		c.Next()
	}
}
