package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: check for access key and verify it
		c.Next()
	}
}
