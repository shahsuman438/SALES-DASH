package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/shahsuman438/SALES-DASH/server/docs"
	"github.com/shahsuman438/SALES-DASH/server/pkg/config"
	"github.com/shahsuman438/SALES-DASH/server/pkg/database"
	"github.com/shahsuman438/SALES-DASH/server/pkg/middleware"
	"github.com/shahsuman438/SALES-DASH/server/pkg/notification"
	"github.com/shahsuman438/SALES-DASH/server/pkg/product"
	"github.com/shahsuman438/SALES-DASH/server/pkg/reports"
	"github.com/shahsuman438/SALES-DASH/server/pkg/sales"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() {
	database.ConnectDB()

	gin.SetMode(config.Cnfg.ServerMode)

	server := gin.New()

	server.Use(gin.Recovery())

	// setup middleware
	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.JSONLoggerMiddleware())

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	notification.StartModule(server)
	product.StartModule(server)
	sales.StartModule(server)
	reports.StartModule(server)

	server.Run(fmt.Sprintf("%s:%s", config.Cnfg.ServerHost, config.Cnfg.ServerPort))
}
