package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/config"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/middleware"
)

func Start() {
	gin.SetMode(config.Cnfg.ServerMode)

	server := gin.New()

	server.Use(gin.Recovery())

	// setup middleware
	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.Auth())
	server.Use(middleware.JSONLoggerMiddleware())

	server.Run(fmt.Sprintf("%s:%s", config.Cnfg.ServerHost, config.Cnfg.ServerPort))
}
