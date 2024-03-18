package reports

import (
	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/response"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/utils/logger"
)

func setupControllers(engine *gin.Engine) {
	r := engine.Group("/reports")
	r.GET("/summery", getSummery)
	r.GET("/sales-by-product", getSalesByProduct)
	r.GET("/sales-by-brand", getSalesByBrand)
}

func getSummery(c *gin.Context) {
	report, err := GetSummeryReports(c)
	if err != nil {
		logger.Error("Error generating summery reports", err)
		response.BadRequest(c, err)
		return
	}
	response.Success(c, report)
}
func getSalesByProduct(c *gin.Context) {
	report, err := GetSalesByProduct(c)
	if err != nil {
		logger.Error("Error generating Sales by Product report", err)
		response.BadRequest(c, err)
		return
	}
	response.Success(c, report)
}
func getSalesByBrand(c *gin.Context) {
	report, err := GetSalesByBrand(c)
	if err != nil {
		logger.Error("Error generating Sales by Brand reports", err)
		response.BadRequest(c, err)
		return
	}
	response.Success(c, report)
}
