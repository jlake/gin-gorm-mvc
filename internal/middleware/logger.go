package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger カスタムロガーミドルウェア
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		log.Printf("[%s] %s %s | %d | %s",
			method,
			path,
			c.ClientIP(),
			statusCode,
			latency,
		)
	}
}
