package reports

import (
	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/response"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/utils/logger"
)

func setupControllers(engine *gin.Engine) {
	r := engine.Group("/reports")
	r.GET("/summery", getSummery)
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
