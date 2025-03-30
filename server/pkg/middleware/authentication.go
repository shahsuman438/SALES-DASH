package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/server/pkg/utils/logger"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !validateToken(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

func validateToken(token string) bool {
	
	request, err := http.NewRequest("POST", "http://localhost:8888/validate", nil)
	if err != nil {
		return false
	}
	request.Header.Set(
		"Authorization", token,
	)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	if response.StatusCode == http.StatusUnauthorized {
		logger.Error("Invalid Credentials", nil)
		return false
	} else if response.StatusCode != http.StatusOK {
		logger.Error("Error calling auth service", nil)
		return false
	}

	var jsonFromService interface{}
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		logger.Error("", err)
		return false
	}

	return true
}
