package sales

import (
	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/server/pkg/response"
)

func setupControllers(engine *gin.Engine) {
	r := engine.Group("/sales")
	r.GET("", GetAllSales)
}

// @Summary Get All Sales
// @Description Get All Sales
// @Produce json
// @Success 200 {array} string "ok"
// @Failure 400 {object} string "error"
// @Router /sales [get]
func GetAllSales(c *gin.Context) {
	data, err := Fetch(c)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	response.Success(c, data)
}
