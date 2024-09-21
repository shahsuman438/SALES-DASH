package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/server/pkg/config"
)

// SetCORSHeaders sets the required CORS headers
func SetCORSHeaders(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", config.Cnfg.AppURL)
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Authorization")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
}

// CORSMiddleware is a CORS middleware for Gin
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		SetCORSHeaders(c)

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	}
}
