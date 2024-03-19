package notification

import "github.com/gin-gonic/gin"

// SetupControllers sets up the necessary routes for the SSE endpoint.
func SetupControllers(router *gin.Engine) {
	router.GET("/events", HandleSSEConnection)
}

// HandleSSEConnection handles the connection to the SSE endpoint.
func HandleSSEConnection(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	StartServerSentEvents(ctx)
}
