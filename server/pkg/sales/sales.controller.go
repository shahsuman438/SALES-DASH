package sales

import (
	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/response"
)

func setupControllers(engine *gin.Engine) {
	r := engine.Group("/sales")
	r.GET("", GetAllSales)
}

func GetAllSales(c *gin.Context) {
	data, err := Fetch(c)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	response.Success(c, data)
}
