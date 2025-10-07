package middleware

import (
	"gin-gorm-mvc/pkg/response"
	"log"

	"github.com/gin-gonic/gin"
)

// Recovery パニックリカバリーミドルウェア
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				response.InternalServerError(c, "Internal server error")
				c.Abort()
			}
		}()
		c.Next()
	}
}
