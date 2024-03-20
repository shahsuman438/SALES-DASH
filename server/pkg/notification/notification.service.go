package notification

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	messageChannel = make(chan string)
)

// StartServerSentEvents sets up the server-sent events endpoint.
func StartServerSentEvents(c *gin.Context) {
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Send an initial message to connect client
	fmt.Fprintf(c.Writer, "data: %s\n\n", `{"message": ""}`)
	flusher.Flush()

	for {
		select {
		case message := <-messageChannel:
			fmt.Fprintf(c.Writer, "data: %s\n\n", `{"message": "`+message+`"}`)
			flusher.Flush()
		case <-c.Writer.CloseNotify():
			// Exit loop if client disconnected
			return
		}
	}
}

// SendMessage sends a message to all connected clients.
func SendMessageToClients(message string) {
	select {
	case messageChannel <- message:
	default:
	}
}
