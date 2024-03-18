package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/config"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/database"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/middleware"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/product"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/reports"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/sales"
)

func Start() {
	database.ConnectDB()

	gin.SetMode(config.Cnfg.ServerMode)

	server := gin.New()

	server.Use(gin.Recovery())

	// setup middleware
	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.JSONLoggerMiddleware())

	product.StartModule(server)
	sales.StartModule(server)
	reports.StartModule(server)

	server.Run(fmt.Sprintf("%s:%s", config.Cnfg.ServerHost, config.Cnfg.ServerPort))
}
