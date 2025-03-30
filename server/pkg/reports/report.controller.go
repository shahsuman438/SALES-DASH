package reports

import (
	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/server/pkg/response"
	loggerservice "github.com/shahsuman438/SALES-DASH/server/pkg/services/logger-service"
	"github.com/shahsuman438/SALES-DASH/server/pkg/utils/logger"
)

func setupControllers(engine *gin.Engine) {
	r := engine.Group("/reports")
	r.GET("/summery", getSummery)
	r.GET("/sales-by-product", getSalesByProduct)
	r.GET("/sales-by-brand", getSalesByBrand)
}

// @Summary Get Summery Reports
// @Description Get Summery Reports
// @Produce json
// @Success 200 {object} SummeryReport "ok"
// @Failure 400 {object} string "error"
// @Router /reports/summery [get]
func getSummery(c *gin.Context) {
	report, err := GetSummeryReports(c)
	if err != nil {
		logger.Error("Error generating summery reports", err)
		response.BadRequest(c, err)
		return
	}
	loggerservice.WriteLog(&loggerservice.LoggerPayload{Name: "APP", Data: "Summery report is generated and responded to APP"})
	response.Success(c, report)
}

// @Summary Get Reports Sales by product
// @Description Get Sales by product reports
// @Produce json
// @Success 200 {array} string "ok"
// @Failure 400 {object} string "error"
// @Router /reports/sales-by-product [get]
func getSalesByProduct(c *gin.Context) {
	report, err := GetSalesByProduct(c)
	if err != nil {
		logger.Error("Error generating Sales by Product report", err)
		response.BadRequest(c, err)
		return
	}
	loggerservice.WriteLog(&loggerservice.LoggerPayload{Name: "APP", Data: "Sales report is generated and responded to APP"})
	response.Success(c, report)
}

// @Summary Get Reports Sales by brand
// @Description Get Sales by product brand
// @Produce json
// @Success 200 {array} string "ok"
// @Failure 400 {object} string "error"
// @Router /reports/sales-by-brand [get]
func getSalesByBrand(c *gin.Context) {
	report, err := GetSalesByBrand(c)
	if err != nil {
		logger.Error("Error generating Sales by Brand reports", err)
		response.BadRequest(c, err)
		return
	}
	loggerservice.WriteLog(&loggerservice.LoggerPayload{Name: "APP", Data: "Sales by brand report is generated and responded to APP"})
	response.Success(c, report)
}
