package middleware

import (
	"evo-blog-gf/internal/known"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-ID")
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Set(known.XRequestIdKey, requestId)
		c.Writer.Header().Set("X-Request-ID", requestId)
	}
}
