package product

import (
	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/response"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/utils/logger"
)

func setupControllers(engine *gin.Engine) {
	r := engine.Group("/product")
	r.GET("", GetAllProducts)
	r.POST("/add", AddProduct)
	r.POST("/addMany", AddManyProduct)
}

func AddProduct(c *gin.Context) {
	var payload Product

	err := c.BindJSON(&payload)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	err = Add(c, &payload)
	if err != nil {
		logger.Error("Unable to Save to Product Error:", err)
		response.BadRequest(c, err)
		return
	}
	response.SuccessfullyCreate(c, "Product Created")
}

func AddManyProduct(c *gin.Context) {
	var payload []Product

	err := c.BindJSON(&payload)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	err = AddMany(c, &payload)
	if err != nil {
		logger.Error("Unable to Save to Product Error:", err)
		response.BadRequest(c, err)
		return
	}
	response.SuccessfullyCreate(c, "Product Created")
}

func GetAllProducts(c *gin.Context) {
	data, err := Fetch(c)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	response.Success(c, data)
}
