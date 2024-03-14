package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type LogEntry struct {
	Method       string `json:"method"`
	Path         string `json:"path"`
	RemoteAddr   string `json:"remoteAddr"`
	ResponseTime string `json:"responseTime"`
	StartTime    string `json:"startTime"`
	StatusCode   int    `json:"statusCode"`
}

func JSONLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		entry := LogEntry{
			Method:       params.Method,
			Path:         params.Path,
			RemoteAddr:   params.ClientIP,
			ResponseTime: params.Latency.String(),
			StartTime:    params.TimeStamp.Format("2006/01/02 - 15:04:05"),
			StatusCode:   params.StatusCode,
		}

		logJSON, err := json.Marshal(entry)
		if err != nil {
			// Handle JSON marshaling error
			return err.Error()
		}

		return string(logJSON) + "\n"
	})
}
